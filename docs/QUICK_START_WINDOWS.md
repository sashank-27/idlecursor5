# Quick Start Guide for Windows

Get the Activity Presence Controller running in minutes on your Windows machine.

## Prerequisites (One-Time Setup)

### 1. Install Python 3.9+
- Download from https://www.python.org/downloads/
- **IMPORTANT**: Check the box "Add Python to PATH" during installation
- Click "Install Now"

### 2. Install Node.js (LTS)
- Download from https://nodejs.org/
- Click "Install" and follow the prompts

### Verify Installation
Open Command Prompt and run:
```
python --version
node --version
npm --version
```

---

## Automatic Setup (Recommended)

### Option A: Full Automatic Setup (Script Handles Everything)
```batch
cd path\to\idlecursor5
full-setup.bat
```
This installs all dependencies and prepares everything. When done, follow the instructions on screen.

### Option B: Interactive Launcher (Menu-Based)
```batch
cd path\to\idlecursor5
launcher.bat
```
Choose from the menu:
1. Full Setup
2. Run Agent Only
3. Run Dashboard Only
4. Run Both in Separate Windows

---

## Manual Setup

### Step 1: Copy Project to Your Machine
- Extract the project to a folder (e.g., `C:\Users\YourName\Desktop\idlecursor5`)

### Step 2: Setup Agent
```batch
cd path\to\idlecursor5\agent
setup-and-run.bat
```
This will:
- Check Python installation
- Install Python dependencies from `requirements.txt`
- Start the agent on `http://127.0.0.1:8787`

### Step 3: Setup Dashboard (In a New Command Prompt)
```batch
cd path\to\idlecursor5\web
npm-setup-and-run.bat
```
This will:
- Check Node.js installation
- Install npm dependencies
- Start the dashboard on `http://localhost:5173`

### Step 4: Open in Browser
Open http://localhost:5173 in your browser and click "Start"

---

## What Each Script Does

| Script | Purpose | Run From |
|--------|---------|----------|
| `full-setup.bat` | One-time setup - installs all dependencies | Project root |
| `launcher.bat` | Interactive menu to start services | Project root |
| `agent/setup-and-run.bat` | Setup and run just the agent | Project root or agent dir |
| `web/npm-setup-and-run.bat` | Setup and run just the dashboard | Project root or web dir |

---

## Troubleshooting

### "Python/Node not found"
- **Solution**: Reinstall from the links above
- Make sure to check "Add to PATH" during installation
- **Close all terminals** after installing and try again

### Port Already in Use
- **Agent on port 8787**: Open Command Prompt and run:
  ```
  netstat -ano | findstr :8787
  taskkill /PID <process_id> /F
  ```
- **Dashboard on port 5173**: Same process but use `:5173`

### "Failed to install dependencies"
- Try manually installing:
  ```
  cd agent
  pip install --upgrade pip
  pip install -r requirements.txt
  ```

### Dashboard shows "Failed to fetch"
- Make sure agent is running (`http://127.0.0.1:8787`)
- Check agent terminal for error messages
- Make sure Windows firewall isn't blocking port 8787

---

## Environment Variables (Optional)

If you need to customize the agent, you can set these before running:

```batch
set APC_BIND=127.0.0.1:8787              REM Listen address and port
set APC_CLOUD_MODE=false                 REM false = cursor movement, true = headless
set APC_ALLOW_INSECURE=true              REM true for HTTP (dev), false for HTTPS
set APC_PAIRING_TOKEN=your-secret        REM Optional: authentication token
```

---

## Running on Another Windows Machine

1. Copy the entire `idlecursor5` folder to the target machine
2. On that machine, open Command Prompt in the folder
3. Run: `full-setup.bat`
4. When done, open `launcher.bat` to start services

---

## First Time Use

1. Open http://localhost:5173 in your browser
2. Click the "Start" button
3. Watch your cursor move automatically (if enabled)
4. Check the logs panel on the right
5. Click "Stop" to stop the agent
6. Use the controls to start/stop as needed

---

## Getting Help

- **Check the logs** in the dashboard (right panel)
- **Check terminal errors** - look at both the agent and dashboard terminal windows
- **Firewall**: Windows Firewall might block ports - temporarily disable to test
- **Antivirus**: Disable temporarily if Python installation is blocked

---

## What's Running

### Agent (Port 8787)
- Python Flask HTTP server
- Manages cursor movement
- Provides REST API for dashboard
- Handles all system-level operations

### Dashboard (Port 5173)
- React web app (PWA)
- Control panel for the agent
- Shows logs and status
- Can be accessed from other machines on the network

---

## Next Steps

- Read [SETUP_WINDOWS.md](SETUP_WINDOWS.md) for detailed manual setup
- Read [START_HERE.md](START_HERE.md) for feature overview
- Check [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md) for advanced configuration

