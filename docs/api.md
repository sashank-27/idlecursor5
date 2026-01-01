# Local Agent API (Draft)

Base: `https://127.0.0.1:8787`

Authentication: pairing token exchanged for session token (header `Authorization: Bearer <token>`). For now, stub endpoints accept unauthenticated requests in development mode.

## Endpoints
- `GET /health` → `{"status":"ok"}`.
- `GET /status` → agent state, mode, active flag, userPresent, policyLock, nextAction, startedAt.
- `POST /session/start` → `{ "mode": "meeting|build|focus|custom", "randomness": 0-1, "idleThresholdSeconds": 60, "maxDurationMinutes": 120 }`.
- `POST /session/stop` → stops simulated activity.
- `POST /policy/lock` → `{ "locked": true|false }`.
- `GET /logs?since=<rfc3339>` → `{ "entries": [ { "ts": "...", "action": "move", "meta": {} } ] }`.
- `GET /stream` (WSS) → pushes state snapshots and user-activity events.

## Status Payload (example)
```json
{
  "state": "active",
  "mode": "meeting",
  "userPresent": false,
  "policyLocked": false,
  "nextAction": "micro-move in 12s",
  "startedAt": "2026-01-01T12:00:00Z"
}
```

## Notes
- TLS: self-signed certificate pinned on first use; configurable via env vars or config file.
- Rate limits: minimal for localhost, but endpoints should be idempotent and reject conflicting sessions.
- Logging: entries stay local; exported only on user action.
