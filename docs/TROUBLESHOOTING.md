# Windows Setup Troubleshooting Guide

Having issues setting up the Activity Presence Controller on Windows? This guide covers common problems and solutions.

## Prerequisites Issues

### Python Installation

**Problem**: `python: command not found` or `'python' is not recognized`

**Solution**:
1. Install Python from https://www.python.org/downloads/
2. During installation, **check the box "Add Python to PATH"**
3. Complete the installation
4. **Restart your terminal/command prompt**
5. Verify: `python --version`

**If still not working**:
- Uninstall Python completely (Control Panel → Programs → Uninstall)
- Reinstall with "Add Python to PATH" checked
- Use full path: `C:\Users\YourName\AppData\Local\Programs\Python\Python311\python.exe --version`

---

### Node.js Installation

**Problem**: `node: command not found` or `'node' is not recognized`

**Solution**:
1. Install Node.js from https://nodejs.org/ (LTS recommended)
2. Complete the installation (it adds to PATH automatically)
3. **Restart your terminal/command prompt**
4. Verify: `node --version` and `npm --version`

**If still not working**:
- Check if Node is in PATH: `echo %PATH%` in Command Prompt
- Uninstall and reinstall Node.js
- Restart your computer

---

## Dependency Installation Issues

### "Failed to Install Requirements"

**Problem**: `setup-and-run.bat` or `pip install` fails

**Solution**:
1. Open Command Prompt as Administrator
2. Navigate to the `agent` directory
3. Run:
   ```batch
   python -m pip install --upgrade pip
   pip install -r requirements.txt
   ```
4. If one package fails, install individually:
   ```batch
   pip install Flask==3.0.0
   pip install Flask-CORS==4.0.0
   pip install python-dotenv==1.0.0
   pip install Werkzeug==3.0.1
   ```

**If caching is the issue**:
```batch
pip cache purge
pip install -r requirements.txt --no-cache-dir
```

---

### "Failed to Install npm Dependencies"

**Problem**: `npm install` fails in web directory

**Solution**:
1. Clear npm cache:
   ```batch
   npm cache clean --force
   ```
2. Try installing again:
   ```batch
   npm install
   ```
3. If still failing, try:
   ```batch
   npm install --verbose
   ```

**Check for disk space**:
- Make sure you have at least 500 MB free
- npm packages can be large

---

## Port/Network Issues

### "Port Already in Use"

**Problem**: `Address already in use` on port 8787 or 5173

**Solution - Find and kill the process**:

For Agent (port 8787):
```batch
netstat -ano | findstr :8787
taskkill /PID <process_id> /F
```

For Dashboard (port 5173):
```batch
netstat -ano | findstr :5173
taskkill /PID <process_id> /F
```

Replace `<process_id>` with the actual number from netstat.

**Alternative - Use different port**:
```batch
set APC_BIND=127.0.0.1:9000
python main.py
```

---

### "Failed to Fetch" in Dashboard

**Problem**: Dashboard shows "Failed to fetch" error

**Solution**:
1. Make sure Agent is running:
   - Check terminal shows: `Running on http://127.0.0.1:8787`
2. Check Windows Firewall:
   - Open Windows Defender Firewall
   - Click "Allow an app through firewall"
   - Ensure Python is allowed or allow port 8787
3. Check agent is responding:
   ```batch
   curl http://127.0.0.1:8787/health
   ```
   or in PowerShell:
   ```powershell
   Invoke-WebRequest -Uri http://127.0.0.1:8787/health -UseBasicParsing
   ```
4. If error persists:
   - Check agent terminal for error messages
   - Make sure dashboard is on `http://localhost:5173` (not 127.0.0.1)

---

### "Connection Refused"

**Problem**: Cannot connect to agent from dashboard

**Solution**:
1. Verify agent is running:
   ```batch
   netstat -ano | findstr :8787
   ```
2. Check if it's listening on 127.0.0.1:
   - Look for `127.0.0.1:8787 LISTENING`
3. Try accessing directly:
   - Open browser to `http://localhost:8787/health`
   - Should see: `{"status":"ok"}`
4. If not working:
   - Restart the agent
   - Check for error messages in the terminal
   - Try different port: `set APC_BIND=0.0.0.0:8787`

---

## Script Execution Issues

### "Cannot be loaded because running scripts is disabled"

**Problem**: PowerShell won't run `.ps1` scripts

**Solution**:
1. Open PowerShell as Administrator
2. Run:
   ```powershell
   Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
   ```
3. Type `Y` and press Enter
4. Try running the script again

Or run with bypass:
```powershell
powershell -ExecutionPolicy Bypass -File full-setup.ps1
```

---

### "Batch file runs but nothing happens"

**Problem**: `setup-and-run.bat` opens and closes immediately

**Solution**:
1. Don't double-click the batch file
2. Open Command Prompt manually:
   - In the agent or web directory, type: `cmd`
3. Then run: `setup-and-run.bat`
4. This way you can see any error messages

Or edit the batch file and add `pause` at the end so it doesn't close.

---

## Runtime Issues

### Cursor Not Moving

**Problem**: Agent runs but cursor doesn't move

**Solution**:
1. Check status in dashboard - should show "Active"
2. Make sure `APC_CLOUD_MODE=false` (check agent terminal output)
3. Try moving mouse while agent is running - it should pause for 15 seconds
4. Check logs in dashboard for `micro_move` entries
5. Give the application focus by clicking on desktop
6. Check for error messages in agent terminal

**On laptops with touchpads**:
- Try with external mouse
- Touchpad sometimes interferes with input simulation

---

### Agent Crashes on Startup

**Problem**: Agent starts but immediately crashes

**Possible cause**: Missing Windows API

**Solution**:
1. Make sure you're on Windows 10 or 11 (Windows 7/8 not supported)
2. Try starting in cloud mode:
   ```batch
   set APC_CLOUD_MODE=true
   python main.py
   ```
3. Check for error message in terminal
4. Share error with support with the full message

---

### Dashboard Won't Load

**Problem**: Browser shows blank page or loading spinner

**Solution**:
1. Check browser console (F12 → Console)
2. Hard refresh: `Ctrl+Shift+R`
3. Try different browser (Chrome, Edge, Firefox)
4. Check dashboard terminal for errors (should show Vite messages)
5. Make sure you're on correct URL: `http://localhost:5173`

---

## Firewall Issues

### Windows Defender Firewall Blocking

**Problem**: "Cannot connect" despite everything running

**Solution**:
1. Open Settings → Network & Internet → Windows Defender Firewall
2. Click "Allow an app through firewall"
3. Click "Change settings"
4. Add Python:
   - Click "Allow another app"
   - Click "Browse"
   - Navigate to Python executable
   - Click "Add"
5. Click OK

Or temporarily disable firewall (not recommended for production):
```batch
netsh advfirewall set allprofiles state off
```

To re-enable:
```batch
netsh advfirewall set allprofiles state on
```

---

## Virtual Machine Issues

### "Platform not implemented" error

**Problem**: Running in a VM and getting platform error

**Solution**:
- Agent requires real Windows with actual hardware
- For VMs, use cloud mode:
  ```batch
  set APC_CLOUD_MODE=true
  python main.py
  ```
- Cloud mode doesn't move cursor (headless)

---

## Antivirus/Security Software

### Installation Blocked

**Problem**: Antivirus blocks Python/Node installation

**Solution**:
1. Temporarily disable antivirus
2. Install Python and Node
3. Re-enable antivirus
4. Run the project

---

## Getting More Help

### Check These Files
1. **Agent terminal** - Look for error messages when starting agent
2. **Dashboard browser console** - F12 → Console tab
3. **Event Viewer** - Windows Event Viewer for system errors
4. **Network issues** - `ping localhost` or `ipconfig`

### Collect Information
When reporting issues, gather:
1. Output from `python --version`
2. Output from `node --version` and `npm --version`
3. Full error message from terminal (copy entire output)
4. Screenshot of the issue
5. What you were doing when it happened

### Common Fixes (Try These First)
1. Restart your computer
2. Close all terminals and open fresh ones
3. Uninstall and reinstall Python/Node
4. Delete `node_modules` and `venv` folders and reinstall
5. Check for Windows updates and install them
6. Disable firewall temporarily to test

