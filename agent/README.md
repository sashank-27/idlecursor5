# Agent Setup Scripts

This directory contains the Python-based Activity Presence Controller agent with automatic setup scripts for Windows.

## Quick Start

### Easiest Method (Automatic)
From this directory, double-click or run:
```batch
setup-and-run.bat
```

This will:
1. Check if Python is installed
2. Install all Python dependencies (Flask, Flask-CORS, etc.)
3. Start the agent on `http://127.0.0.1:8787`

### What's Running
- **Agent**: `main.py` - Python Flask HTTP server
- **Port**: 8787 (configurable via `APC_BIND` environment variable)
- **API**: REST API for the dashboard

## Files

### Scripts
- **setup-and-run.bat** - Batch script for automatic setup and run (recommended for cmd.exe)
- **setup-and-run.ps1** - PowerShell script for automatic setup and run

### Python Files
- **main.py** - Entry point, reads environment variables and starts Flask server
- **internal/api/server.py** - Flask HTTP server with all endpoints
- **internal/behavior/engine.py** - Core behavior engine (cursor movement logic)
- **internal/system/platform.py** - Windows API bindings (Win32 calls)

### Configuration
- **requirements.txt** - Python package dependencies

## Environment Variables

You can customize the agent by setting these before running:

| Variable | Default | Description |
|----------|---------|-------------|
| `APC_BIND` | `0.0.0.0:8787` | Server address and port |
| `APC_CLOUD_MODE` | `false` | true=no cursor, false=cursor movement |
| `APC_ALLOW_INSECURE` | `false` | true=HTTP only (dev), false=HTTPS required |
| `APC_PAIRING_TOKEN` | (none) | Optional bearer token for auth |
| `APC_CERT_FILE` | (none) | Path to TLS cert |
| `APC_KEY_FILE` | (none) | Path to TLS key |

Example:
```batch
set APC_BIND=127.0.0.1:9000
set APC_CLOUD_MODE=false
python main.py
```

## API Endpoints

- `GET /health` - Health check
- `GET /status` - Current state
- `POST /session/start` - Start activity session
- `POST /session/stop` - Stop activity session
- `POST /policy/lock` - Lock/unlock policy
- `GET /logs` - Get action logs
- `GET /stream` - Server-sent events stream

## Troubleshooting

### Python Not Found
- Install Python 3.9+ from https://www.python.org/downloads/
- **Make sure to check "Add Python to PATH"**

### Port Already in Use
```batch
netstat -ano | findstr :8787
taskkill /PID <pid> /F
```

### Dependency Installation Failed
Try manually:
```batch
pip install --upgrade pip
pip install -r requirements.txt
```

### Still Having Issues?
- Check that Python is in PATH: `python --version`
- Check that all dependencies installed: `pip list`
- Look at error messages in the terminal
- Check Windows Firewall settings

## Running Elsewhere

The agent works on any Windows machine with Python 3.9+ installed. Just copy this folder and run the setup script.
