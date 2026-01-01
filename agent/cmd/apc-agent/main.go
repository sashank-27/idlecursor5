package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"apc/internal/api"
)

func main() {
	cfg := api.Config{
		Addr:          envOrDefault("APC_BIND", "0.0.0.0:8787"),
		PairingToken:  os.Getenv("APC_PAIRING_TOKEN"),
		CertFile:      os.Getenv("APC_CERT_FILE"),
		KeyFile:       os.Getenv("APC_KEY_FILE"),
		AllowInsecure: os.Getenv("APC_ALLOW_INSECURE") == "true",
		CloudMode:     os.Getenv("APC_CLOUD_MODE") == "true",
	}
	srv := api.NewServer(cfg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("agent listening on %s (https if certs set)\n", cfg.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutdown requested; draining...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}

func envOrDefault(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
