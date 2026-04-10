# Setup Guide for Windows Machines

This guide will help you set up and run the Activity Presence Controller on a fresh Windows machine.

## Prerequisites

### Option 1: Automatic Setup (Recommended)

We provide two automated setup scripts that handle everything:

#### PowerShell Script (recommended)
```powershell
# 1. Open PowerShell as Administrator
# 2. Navigate to the agent directory
cd path\to\idlecursor5\agent

# 3. Run the setup script
powershell -ExecutionPolicy Bypass -File setup-and-run.ps1
```

#### Batch Script (simple alternative)
```batch
# 1. Open Command Prompt (cmd.exe)
# 2. Navigate to the agent directory
cd path\to\idlecursor5\agent

# 3. Run the setup script
setup-and-run.bat
```

---

## Manual Setup (If Scripts Don't Work)

### Step 1: Install Python

1. Download Python 3.9+ from https://www.python.org/downloads/
2. **IMPORTANT**: During installation, check the box **"Add Python to PATH"**
3. Click "Install Now"
4. Verify installation by opening Command Prompt and running:
   ```
   python --version
   ```

### Step 2: Install Node.js (for dashboard)

1. Download Node.js from https://nodejs.org/ (LTS version recommended)
2. Run the installer and follow prompts
3. Verify installation:
   ```
   node --version
   npm --version
   ```

### Step 3: Install Agent Dependencies

```batch
cd path\to\idlecursor5\agent
pip install -r requirements.txt
```

### Step 4: Install Dashboard Dependencies

```batch
cd path\to\idlecursor5\web
npm install
```

### Step 5: Run Both Services

**Terminal 1 - Agent:**
```batch
cd path\to\idlecursor5\agent
set APC_ALLOW_INSECURE=true
set APC_CLOUD_MODE=false
set APC_BIND=127.0.0.1:8787
python main.py
```

**Terminal 2 - Dashboard:**
```batch
cd path\to\idlecursor5\web
npm run dev
```

Open http://localhost:5173 in your browser.

---

## Troubleshooting

### "Python is not installed or not in PATH"
- Reinstall Python from https://www.python.org/downloads/
- **Make sure to check "Add Python to PATH"** during installation
- Close any open terminals and try again

### "pip: command not found"
- Ensure Python was installed with pip
- Try: `python -m pip install -r requirements.txt`

### "Node/npm not found"
- Install Node.js from https://nodejs.org/
- Close terminals and try again

### "Port 8787 already in use"
- Kill the previous agent process:
  ```
  taskkill /IM python.exe /F
  ```
- Or use a different port: `set APC_BIND=127.0.0.1:8888`

### Dashboard shows "Failed to fetch"
- Make sure the agent is running on http://127.0.0.1:8787
- Check that no firewall is blocking the connection
- Look at the agent terminal for error messages

---

## Environment Variables

You can customize the agent by setting these environment variables before running:

```batch
set APC_BIND=127.0.0.1:8787              REM Server address and port
set APC_CLOUD_MODE=false                 REM false for cursor movement, true for headless
set APC_ALLOW_INSECURE=true              REM true for HTTP dev mode
set APC_PAIRING_TOKEN=your-secret        REM Optional: token for auth
set APC_CERT_FILE=path/to/cert.pem       REM Optional: TLS certificate
set APC_KEY_FILE=path/to/key.pem         REM Optional: TLS key
```

---

## First Time Running

1. Run the setup script (automatic or manual)
2. Open http://localhost:5173 in your browser
3. Click "Start" button
4. Watch your cursor move automatically (if enabled)
5. Check the logs on the right side
6. Click "Stop" to stop the agent

---

## Getting Help

- Check the logs in the dashboard for errors
- Make sure both agent and dashboard terminals show no errors
- Verify Python and Node.js versions are recent enough
- Check that ports 8787 (agent) and 5173 (dashboard) are available

