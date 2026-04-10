# Setup Scripts Reference

Quick reference guide for all setup and run scripts.

## Script Hierarchy

```
idlecursor5/ (project root)
├── full-setup.bat ..................... Full setup (agent + dashboard)
├── full-setup.ps1 ..................... Full setup (PowerShell)
├── launcher.bat ....................... Interactive menu launcher
│
├── agent/
│   ├── setup-and-run.bat .............. Setup + run agent
│   ├── setup-and-run.ps1 .............. Setup + run agent (PowerShell)
│   └── main.py ........................ Entry point
│
└── web/
    ├── npm-setup-and-run.bat .......... Setup + run dashboard
    └── src/App.tsx .................... Entry point
```

## Decision Tree

### First Time Setup?

```
START
 │
 ├─→ Want easiest method?
 │   ├─→ YES: Run full-setup.bat from project root
 │   │   └─→ Then run launcher.bat
 │   └─→ NO: Read manual setup guide
 │
 ├─→ Using Command Prompt (cmd.exe)?
 │   ├─→ YES: Use .bat scripts
 │   └─→ NO: Use .ps1 scripts
 │
 ├─→ Platform Windows?
 │   ├─→ YES: Use setup scripts
 │   └─→ NO: Use manual setup
 │
└─→ DONE: Open http://localhost:5173
```

## Commands by Goal

### "I just want to get it running now"
```batch
cd path\to\idlecursor5
full-setup.bat                    # One-time setup (installs all deps)
launcher.bat                       # Run with menu
```

### "I want a menu to choose what to run"
```batch
cd path\to\idlecursor5
launcher.bat                       # Shows menu with options
```

### "I want to run just the agent"
```batch
cd path\to\idlecursor5\agent
setup-and-run.bat
```

### "I want to run just the dashboard"
```batch
cd path\to\idlecursor5\web
npm-setup-and-run.bat
```

### "I prefer manual control"

Agent:
```batch
cd path\to\idlecursor5\agent
pip install -r requirements.txt
set APC_ALLOW_INSECURE=true
set APC_CLOUD_MODE=false
python main.py
```

Dashboard (new terminal):
```batch
cd path\to\idlecursor5\web
npm install
npm run dev
```

---

## Script Comparison

| Script | Purpose | When to Use | Requirements |
|--------|---------|------------|--------------|
| `full-setup.bat` | Install everything once | First time only | Admin CMD |
| `launcher.bat` | Interactive menu | Any time | Agent + Dashboard ready |
| `setup-and-run.bat` (agent) | Install + run agent | Per session | Python installed |
| `setup-and-run.ps1` (agent) | Install + run agent (PS) | Per session | Python + PowerShell |
| `npm-setup-and-run.bat` | Install + run dashboard | Per session | Node/npm installed |

---

## What Each Script Does

### `full-setup.bat` (Run once)
```
1. Check Python installation
2. Navigate to agent/
3. Install Python deps (pip install -r requirements.txt)
4. Navigate to web/
5. Install npm deps (npm install)
6. Done - shows instructions
```

### `launcher.bat` (Run anytime after full-setup)
```
Display menu:
  1. Full Setup (repeats full-setup)
  2. Run Agent Only
  3. Run Dashboard Only
  4. Run Both (opens 2 new windows)
  5. Exit
```

### `setup-and-run.bat` (Run agent)
```
1. Check Python installed
2. Install pip deps (if not present)
3. Set environment variables
4. Run python main.py
5. Listening on http://127.0.0.1:8787
```

### `npm-setup-and-run.bat` (Run dashboard)
```
1. Check Node.js installed
2. Install npm deps (if not present)
3. Run npm run dev
4. Listening on http://localhost:5173
```

---

## Best Practices

### First Time Setup
1. Run `full-setup.bat` once to install everything
2. Then use `launcher.bat` for easy running

### Regular Use
- Option A: Use `launcher.bat` (menu-based)
- Option B: Open command prompts for agent and dashboard separately

### Development
- Use individual `setup-and-run.bat` scripts
- Allows independent restart of agent/dashboard
- Easier to see individual terminal output

### Troubleshooting
- Run individual setup scripts to see detailed error messages
- Check `node_modules/` and `venv/` folders exist
- Delete and reinstall if dependencies corrupted

---

## Environment Variables Reference

Set before running agent `main.py`:

```batch
# Use localhost address and custom port
set APC_BIND=127.0.0.1:8787

# false = cursor movement on Windows
# true = no cursor (headless/cloud mode)
set APC_CLOUD_MODE=false

# true = HTTP (dev), false = HTTPS required
set APC_ALLOW_INSECURE=true

# Optional: Bearer token authentication
set APC_PAIRING_TOKEN=my-secret-token

# Optional: TLS certificate
set APC_CERT_FILE=C:\path\to\cert.pem

# Optional: TLS key
set APC_KEY_FILE=C:\path\to\key.pem
```

Example with custom variables:
```batch
set APC_BIND=0.0.0.0:9000
set APC_CLOUD_MODE=true
python main.py
```

---

## Port Reference

| Service | Default Port | Configurable |
|---------|-------------|--------------|
| Agent (Python Flask) | 8787 | Yes (APC_BIND) |
| Dashboard (Vite dev) | 5173 | No (see vite.config.ts) |

Check if ports available:
```batch
netstat -ano | findstr :8787
netstat -ano | findstr :5173
```

Kill if in use:
```batch
taskkill /PID <pid> /F
```

---

## Typical Workflow

### Session 1 (One-time setup)
```
1. Extract project to C:\Users\You\Desktop\idlecursor5
2. Open CMD in that directory
3. Run: full-setup.bat
4. Wait for completion
5. Close CMD
```

### Session 2+ (Regular use)
```
1. Open CMD in C:\Users\You\Desktop\idlecursor5
2. Run: launcher.bat
3. Choose option 4 (Run Both)
4. Two windows open with agent + dashboard
5. Browser opens http://localhost:5173
6. Click Start in dashboard
7. Done!
```

---

## Next Steps

- [QUICK_START_WINDOWS.md](QUICK_START_WINDOWS.md) - Beginner guide
- [SETUP_VERIFICATION.md](SETUP_VERIFICATION.md) - Verification checklist
- [TROUBLESHOOTING.md](TROUBLESHOOTING.md) - Fix common issues
- [START_HERE.md](START_HERE.md) - Feature overview

