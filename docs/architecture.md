# Architecture Overview

## Components
- **Local Agent (Go)**: localhost HTTPS/JSON + WSS API. Runs as service; owns behavior engine, OS-level hooks, power inhibitors, and logging. No telemetry.
- **Web Dashboard (React/Vite PWA)**: controls sessions, schedules, and policy locks. Works offline; pairs to agent via token. One-click emergency stop.
- **Behavior Engine**: generates human-like micro-movements/keystrokes, honors idle thresholds, pauses when user activity detected, and caps daily runtime.
- **Optional HID Dongle**: microcontroller firmware acting as standard HID mouse/keyboard for locked-down environments.

## Trust & Security Model
- Localhost-only surface; self-signed cert pinned on first use (TOFU) or user-provided cert.
- Pairing token → short-lived session token; CSRF-protected; same-origin policy bound to localhost.
- No cloud dependency; offline-first. Logs stored locally and exportable by the user.
- Enterprise policy lock: explicit flag to hard-disable during secure sessions; optional allow/deny application list.
- Visible status indicator while active; cannot run silently.

## Data Flows
1) Dashboard issues commands to agent over HTTPS (`/session/start`, `/session/stop`, `/schedule`, `/policy/lock`).
2) Agent streams status via WSS (`/stream`) and exposes recent logs (`/logs`).
3) Behavior engine emits events → logger → log store (local file/SQLite). No external transmission.

## Mode Semantics
- **Meeting Mode**: Keeps presence without cursor drift; minimal key taps or wake signals.
- **Build Mode**: Prevents sleep/screensaver; periodic wake signals; low/no pointer movement.
- **Focus Mode**: Prevents lock only; no pointer moves/keys; uses power inhibitors.

## Platform Hooks (planned)
- **Windows**: `SendInput`, `SetThreadExecutionState`, raw input for detection, service via SCM.
- **macOS**: `CGEventPost`, `IOPMAssertionCreateWithName`, Accessibility permission flow, launchd agent.
- **Linux**: `uinput` virtual devices; Wayland portals where available; `systemd-inhibit` for wake; fallback to HID dongle.

## Scheduling & Rules
- Time windows with recurrence (weekday/weekend), max active per day, and idle-based stop.
- Context-aware: pause on real user input; resume after idle threshold with jitter.
- Emergency stop is always available and OS-global.

## Observability
- Local log of simulated actions (timestamp, action type, parameters, mode).
- Health endpoint for readiness; status stream includes permissions state and next scheduled action.

## Packaging (targets)
- Windows: MSI + service installer.
- macOS: pkg + launchd plist.
- Linux: AppImage/deb/rpm + systemd user service.
- Firmware: prebuilt UF2/hex for common microcontrollers.
