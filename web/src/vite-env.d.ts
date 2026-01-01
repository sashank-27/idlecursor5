/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_AGENT_ORIGIN?: string;  // e.g., "https://apc-agent.fly.dev" or "http://127.0.0.1:8787"
  readonly VITE_AGENT_TOKEN?: string;   // Optional Bearer token for cloud agents
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
