# Setup Verification Checklist

Use this checklist to verify that everything is installed and working correctly.

## Prerequisites Verification

### Python Installation
- [ ] Python 3.9+ installed
  ```batch
  python --version
  ```
  Expected output: `Python 3.9.x` or higher

- [ ] Python in PATH
  ```batch
  where python
  ```
  Expected output: `C:\...\python.exe`

- [ ] pip working
  ```batch
  pip --version
  ```
  Expected output: `pip X.X.X from ...`

### Node.js Installation
- [ ] Node.js installed
  ```batch
  node --version
  ```
  Expected output: `vX.XX.X`

- [ ] npm installed
  ```batch
  npm --version
  ```
  Expected output: `X.X.X`

- [ ] Node in PATH
  ```batch
  where node
  ```
  Expected output: `C:\...\node.exe`

---

## Dependencies Verification

### Agent Dependencies
- [ ] Navigate to agent directory
  ```batch
  cd agent
  ```

- [ ] Verify requirements.txt exists
  ```batch
  dir requirements.txt
  ```

- [ ] Install dependencies
  ```batch
  pip install -r requirements.txt
  ```

- [ ] Verify Flask installed
  ```batch
  python -c "import flask; print(flask.__version__)"
  ```
  Expected output: `3.0.0` or similar

- [ ] Verify Flask-CORS installed
  ```batch
  python -c "import flask_cors; print('OK')"
  ```
  Expected output: `OK`

### Dashboard Dependencies
- [ ] Navigate to web directory
  ```batch
  cd web
  ```

- [ ] Verify package.json exists
  ```batch
  dir package.json
  ```

- [ ] Install npm packages
  ```batch
  npm install
  ```

- [ ] Verify packages installed
  ```batch
  dir node_modules
  ```
  Should show many folders

---

## Agent Verification

### Agent Startup
- [ ] Navigate to agent directory
  ```batch
  cd agent
  ```

- [ ] Set environment variables
  ```batch
  set APC_ALLOW_INSECURE=true
  set APC_CLOUD_MODE=false
  set APC_BIND=127.0.0.1:8787
  ```

- [ ] Start agent
  ```batch
  python main.py
  ```

- [ ] Verify startup message
  Expected output in terminal:
  ```
  agent listening on 127.0.0.1:8787
  Running on http://127.0.0.1:8787
  ```

### Agent API Verification
- [ ] Open new command prompt
- [ ] Check health endpoint
  ```batch
  curl http://127.0.0.1:8787/health
  ```
  Expected output: `{"status":"ok"}`

- [ ] Check status endpoint (may need auth token)
  ```batch
  curl http://127.0.0.1:8787/status
  ```
  Expected output: JSON with state information

- [ ] Port is listening
  ```batch
  netstat -ano | findstr :8787
  ```
  Expected: Should see a LISTENING entry

---

## Dashboard Verification

### Dashboard Startup
- [ ] Navigate to web directory
  ```batch
  cd web
  ```

- [ ] Start development server
  ```batch
  npm run dev
  ```

- [ ] Verify startup message
  Expected output:
  ```
  VITE ready in XXX ms
  ➜ Local: http://localhost:5173/
  ```

### Dashboard Access
- [ ] Open browser
- [ ] Navigate to http://localhost:5173
- [ ] Page loads without errors
- [ ] See control panel with "Start" button
- [ ] Browser console has no major errors (F12)

---

## End-to-End Testing

### Session Test
1. [ ] Both services running (agent on 8787, dashboard on 5173)
2. [ ] Dashboard loaded in browser
3. [ ] Click "Start" button
4. [ ] Dashboard status changes to "Active"
5. [ ] Dashboard shows "micro_move" logs
6. [ ] (If not cloud mode) See cursor moving
7. [ ] Click "Stop" button
8. [ ] Status changes back to "Idle"
9. [ ] Cursor stops moving

### Log Verification
- [ ] Dashboard shows activity logs
- [ ] Logs include timestamps
- [ ] Logs show "micro_move" entries with dx/dy values
- [ ] When you move mouse, see "user_active_pause" log

### Cursor Movement (if not cloud mode)
- [ ] Cursor moves smoothly
- [ ] Movement is random (-100 to +100 pixels)
- [ ] Movement pauses when you move mouse
- [ ] Pause lasts approximately 15 seconds
- [ ] Resume after pause period

---

## Troubleshooting Checklist

If something doesn't work:

Operating System
- [ ] Windows 10 or 11
- [ ] Run as Administrator (if permissions issues)

Network/Firewall
- [ ] Windows Firewall allows Python on port 8787
- [ ] Antivirus not blocking connections
- [ ] Try disabling firewall temporarily to test

Path/Environment
- [ ] Check user PATH includes Python and Node
- [ ] No spaces in installation paths (if possible)
- [ ] Run from command prompt, not file explorer

Terminal Issues
- [ ] Close all terminals, open fresh ones
- [ ] Try both Command Prompt and PowerShell
- [ ] Use "Run as Administrator"

Installation Issues
- [ ] Reinstall Python with PATH option
- [ ] Reinstall Node.js
- [ ] Clear pip cache: `pip cache purge`
- [ ] Clear npm cache: `npm cache clean --force`

---

## Success Criteria

Everything is working when:

✓ Python installed and in PATH
✓ Node/npm installed and in PATH
✓ Agent dependencies installed
✓ Dashboard dependencies installed
✓ Agent runs without errors on port 8787
✓ Agent responds to health check
✓ Dashboard runs without errors on port 5173
✓ Dashboard loads in browser
✓ Dashboard connects to agent
✓ Can start/stop sessions
✓ Logs appear in dashboard
✓ (If enabled) Cursor moves automatically

---

## Next Steps

Once verification complete:

1. Read [QUICK_START_WINDOWS.md](QUICK_START_WINDOWS.md) for usage guide
2. Check [START_HERE.md](START_HERE.md) for feature overview
3. See [docs/api.md](docs/api.md) for API documentation
4. Read [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md) for advanced setup

