package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"apc/internal/behavior"
	"apc/internal/system"
)

// Config wires runtime options for the agent's HTTP server.
type Config struct {
	Addr          string
	PairingToken  string
	CertFile      string
	KeyFile       string
	AllowInsecure bool
	CloudMode     bool // if true, disables local cursor movement on cloud VMs
}

type state struct {
	State        string    `json:"state"`
	Mode         string    `json:"mode"`
	UserPresent  bool      `json:"userPresent"`
	PolicyLocked bool      `json:"policyLocked"`
	NextAction   string    `json:"nextAction"`
	StartedAt    time.Time `json:"startedAt"`
}

type logEntry struct {
	Timestamp time.Time         `json:"ts"`
	Action    string            `json:"action"`
	Meta      map[string]string `json:"meta"`
}

type Server struct {
	addr          string
	pairingToken  string
	allowInsecure bool
	certFile      string
	keyFile       string
	cloudMode     bool

	mu            sync.Mutex
	st            state
	logs          []logEntry
	engine        *behavior.Engine
	sessionCancel context.CancelFunc
	http          *http.Server
}

func NewServer(cfg Config) *Server {
	var plat system.Platform
	var err error

	if cfg.CloudMode {
		log.Println("running in cloud mode (cursor movement disabled)")
		plat = nil // headless mode, no platform
	} else {
		plat, err = system.NewPlatform()
		if err != nil {
			log.Printf("platform init warning: %v", err)
		}
	}

	s := &Server{
		addr:          cfg.Addr,
		pairingToken:  strings.TrimSpace(cfg.PairingToken),
		allowInsecure: cfg.AllowInsecure,
		certFile:      cfg.CertFile,
		keyFile:       cfg.KeyFile,
		cloudMode:     cfg.CloudMode,
		engine:        behavior.New(plat),
	}
	s.st = state{State: "idle", Mode: "", UserPresent: false, PolicyLocked: false, NextAction: "", StartedAt: time.Time{}}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth) // keep unauthenticated for liveness
	mux.HandleFunc("/status", s.cors(s.auth(s.handleStatus)))
	mux.HandleFunc("/session/start", s.cors(s.auth(s.handleStartSession)))
	mux.HandleFunc("/session/stop", s.cors(s.auth(s.handleStopSession)))
	mux.HandleFunc("/policy/lock", s.cors(s.auth(s.handlePolicyLock)))
	mux.HandleFunc("/logs", s.cors(s.auth(s.handleLogs)))
	mux.HandleFunc("/stream", s.cors(s.auth(s.handleStream)))

	s.http = &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}
	return s
}

func (s *Server) ListenAndServe() error {
	if s.certFile != "" && s.keyFile != "" {
		log.Printf("serving HTTPS with cert %s", s.certFile)
		return s.http.ListenAndServeTLS(s.certFile, s.keyFile)
	}

	if !s.allowInsecure {
		return fmt.Errorf("TLS not configured; set APC_CERT_FILE/APC_KEY_FILE or APC_ALLOW_INSECURE=true for dev HTTP")
	}

	log.Printf("WARNING: serving HTTP without TLS on %s (dev mode only)", s.addr)
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleStatus(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	writeJSON(w, http.StatusOK, s.st)
}

func (s *Server) handleStartSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Mode                 string  `json:"mode"`
		Randomness           float64 `json:"randomness"`
		IdleThresholdSeconds int     `json:"idleThresholdSeconds"`
		MaxDurationMinutes   int     `json:"maxDurationMinutes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}

	s.mu.Lock()
	if s.st.State == "active" {
		s.mu.Unlock()
		writeJSON(w, http.StatusConflict, map[string]string{"error": "session already active"})
		return
	}
	if s.st.PolicyLocked {
		s.mu.Unlock()
		writeJSON(w, http.StatusForbidden, map[string]string{"error": "policy locked"})
		return
	}

	sessionCtx, cancel := context.WithCancel(context.Background())
	s.sessionCancel = cancel

	s.st.State = "active"
	s.st.Mode = req.Mode
	s.st.StartedAt = time.Now().UTC()
	s.st.NextAction = "simulate in 5s"
	s.mu.Unlock()

	s.logAction("session_start", map[string]string{"mode": req.Mode})
	go s.engine.Run(sessionCtx, req.Mode, s.logAction)

	writeJSON(w, http.StatusOK, map[string]string{"status": "started"})
}

func (s *Server) handleStopSession(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	s.mu.Lock()
	cancel := s.sessionCancel
	s.sessionCancel = nil
	s.st.State = "idle"
	s.st.NextAction = ""
	s.mu.Unlock()

	if cancel != nil {
		cancel()
	}
	s.logAction("session_stop", map[string]string{})

	writeJSON(w, http.StatusOK, map[string]string{"status": "stopped"})
}

func (s *Server) handlePolicyLock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Locked bool `json:"locked"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid payload"})
		return
	}

	s.mu.Lock()
	s.st.PolicyLocked = req.Locked
	action := "policy_unlocked"
	if req.Locked {
		action = "policy_locked"
		s.st.State = "paused"
	}
	s.mu.Unlock()

	s.logAction(action, map[string]string{})

	writeJSON(w, http.StatusOK, map[string]string{"status": action})
}

func (s *Server) handleLogs(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// TODO: respect ?since filter; this is a stub.
	resp := struct {
		Entries []logEntry `json:"entries"`
	}{Entries: s.logs}
	writeJSON(w, http.StatusOK, resp)
}

func (s *Server) handleStream(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// send initial snapshot
	s.writeEvent(w, "status", s.snapshot())
	flusher.Flush()

	for {
		select {
		case <-r.Context().Done():
			return
		case <-ticker.C:
			s.writeEvent(w, "status", s.snapshot())
			flusher.Flush()
		}
	}
}

func (s *Server) auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.pairingToken == "" {
			next(w, r)
			return
		}

		header := r.Header.Get("Authorization")
		if header == "Bearer "+s.pairingToken {
			next(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
	}
}

func (s *Server) cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Echo caller origin when present; fallback to wildcard for non-browser clients.
		// For stricter control, set APC_ALLOWED_ORIGINS="https://foo.app,https://bar.app".
		if allowList := strings.TrimSpace(os.Getenv("APC_ALLOWED_ORIGINS")); allowList != "" {
			for _, o := range strings.Split(allowList, ",") {
				if strings.EqualFold(strings.TrimSpace(o), origin) {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
		} else if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "false")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func (s *Server) logAction(action string, meta map[string]string) {
	// meta may be nil; keep map to avoid nil in JSON.
	if meta == nil {
		meta = map[string]string{}
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	entry := logEntry{Timestamp: time.Now().UTC(), Action: action, Meta: meta}
	s.logs = append(s.logs, entry)
}

func (s *Server) snapshot() state {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.st
}

func (s *Server) writeEvent(w http.ResponseWriter, event string, payload interface{}) {
	b, err := json.Marshal(payload)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, b)
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON error: %v", err)
	}
}
