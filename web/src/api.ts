export type AgentStatus = {
  state: string;
  mode: string;
  userPresent: boolean;
  policyLocked: boolean;
  nextAction: string;
  startedAt: string;
};

export type AgentLogEntry = {
  ts: string;
  action: string;
  meta: Record<string, string>;
};

const AGENT_BASE = (import.meta.env.VITE_AGENT_ORIGIN as string) || "http://127.0.0.1:8787";
const AGENT_TOKEN = (import.meta.env.VITE_AGENT_TOKEN as string) || "";

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
    ...(init?.headers as Record<string, string> || {})
  };
  
  // Add Bearer token if provided
  if (AGENT_TOKEN) {
    headers["Authorization"] = `Bearer ${AGENT_TOKEN}`;
  }
  
  const res = await fetch(`${AGENT_BASE}${path}`, {
    ...init,
    headers
  });
  if (!res.ok) {
    const text = await res.text();
    throw new Error(text || res.statusText);
  }
  return res.json();
}

export function getStatus() {
  return request<AgentStatus>("/status");
}

export function startSession(body: {
  mode: string;
  randomness: number;
  idleThresholdSeconds: number;
  maxDurationMinutes: number;
}) {
  return request<{ status: string }>("/session/start", {
    method: "POST",
    body: JSON.stringify(body)
  });
}

export function stopSession() {
  return request<{ status: string }>("/session/stop", { method: "POST" });
}

export function setPolicyLock(locked: boolean) {
  return request<{ status: string }>("/policy/lock", {
    method: "POST",
    body: JSON.stringify({ locked })
  });
}

export function getLogs() {
  return request<{ entries: AgentLogEntry[] }>("/logs");
}

export function openStatusStream(onStatus: (s: AgentStatus) => void): () => void {
  const source = new EventSource(`${AGENT_BASE}/stream`, { withCredentials: false });

  const handler = (event: MessageEvent) => {
    try {
      const data = JSON.parse(event.data) as AgentStatus;
      onStatus(data);
    } catch (err) {
      console.error("failed to parse status event", err);
    }
  };

  source.addEventListener("status", handler as EventListener);
  source.onerror = () => {
    // Let the browser retry; minimal handling for now.
  };

  return () => {
    source.removeEventListener("status", handler as EventListener);
    source.close();
  };
}
