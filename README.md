# Activity Presence Controller

An ethical, transparent, cross-platform alternative to browser-only mouse jigglers. The project pairs a localhost companion agent with a PWA dashboard to simulate user presence responsibly, keep systems awake, and provide auditable logs.

## Core Ideas
- Local-first: all control and logs stay on the device; no telemetry.
- Agent + Dashboard: the browser UI can be closed after issuing commands; the agent keeps running.
- OS-level presence: leverages platform APIs (and optional HID hardware) so desktop apps recognize activity.
- Security & consent: explicit enablement, visible indicators, enterprise policy locks, and clear action logs.

## Repository Layout
- `agent/` – Go companion service (localhost API, behavior engine, OS hooks).
- `web/` – React/Vite PWA dashboard for control, scheduling, and monitoring.
- `docs/` – Architecture, API contract, design notes.
- `firmware/` – Placeholder for HID USB firmware reference.
- `scripts/` – Helper scripts (packaging/build stubs).

## Features (planned)
- Modes: Meeting, Build, Focus, plus custom profiles.
- Behavior engine: human-like micro-movements, idle thresholds, context-aware pauses, adaptive jitter.
- Scheduling: recurring windows, caps, emergency stop always available.
- Compliance: consent gating, enterprise policy lock, local logs of simulated actions.
- Cross-platform: Windows (SendInput + ES), macOS (CGEvent + IOPM), Linux (uinput/portals), optional HID dongle.

## Quickstart

### Local Machine (Development & WFH Use)

1) **Agent**: install Go 1.21+, then `cd agent && go build ./cmd/apc-agent`.
2) **Run locally**:
   ```bash
   set APC_ALLOW_INSECURE=true
   set APC_CLOUD_MODE=false
   set APC_BIND=127.0.0.1:8787
   ./apc-agent
   ```
3) **Dashboard**: install Node 20+, then `cd web && npm install && npm run dev`.
4) Open the dashboard at `http://localhost:5173` and start a session. Cursor will move on your actual machine.

### Cloud Deployment (Remote Servers)

For hosting on Fly.io, Render, Railway, or DigitalOcean, see [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md).

**Key difference**: Cloud mode (`APC_CLOUD_MODE=true`) disables cursor movement since cloud servers are headless.
Useful for keeping remote desktop sessions alive or centralized monitoring.

## Security & Privacy Defaults
- No data collection; everything stored locally.
- Localhost-only API with token-based pairing and TLS (TOFU) planned.
- Visible status indicator when active; cannot run silently.
- Policy lock lets enterprises hard-disable during secure sessions.

## Next Steps
- Implement TLS bootstrap and pairing flow for the agent.
- Flesh out behavior engine with OS-specific shims.
- Build dashboard scheduling UI and WebSocket status stream.
- Add firmware reference for HID USB mode.
