# Getting Started Guide

## Quick Navigation

- **Local WFH Setup** → [Run Locally](#run-locally)
- **Cloud Agent + Vercel Dashboard** → [Cloud + Vercel](#cloud--vercel-setup)
- **Advanced Deployments** → See [DEPLOYMENT_CLOUD.md](DEPLOYMENT_CLOUD.md) & [DEPLOYMENT_VERCEL.md](DEPLOYMENT_VERCEL.md)

---

## Run Locally

Keep your machine "active" in Teams/Slack/Zoom. The agent moves your cursor automatically.

### Prerequisites

- Windows 10/11 (macOS/Linux not yet supported)
- Go 1.21+ ([download](https://go.dev/dl/))
- Node 18+ ([download](https://nodejs.org/))

### Steps

1. **Open two terminal windows**

2. **Terminal 1: Start the agent**
   ```powershell
   cd agent
   $env:APC_ALLOW_INSECURE="true"
   $env:APC_CLOUD_MODE="false"
   $env:APC_BIND="127.0.0.1:8787"
   
   go mod tidy
   go run ./cmd/apc-agent/main.go
   ```
   
   You should see:
   ```
   2026/01/01 12:00:00 platform init success: Windows
   2026/01/01 12:00:00 listening on http://127.0.0.1:8787
   ```

3. **Terminal 2: Start the dashboard**
   ```powershell
   cd web
   npm install
   npm run dev
   ```
   
   You should see:
   ```
   VITE v5.0.0  ready in 145 ms
   ➜  Local:   http://localhost:5173/
   ```

4. **Open browser**: Visit `http://localhost:5173`

5. **Start a session**: Click "Start" button
   - Cursor will move every 500ms
   - Cursor pauses for 15 seconds when you move the mouse
   - Check logs on the right to see activity

6. **Stop**: Click "Stop" button

### Verify It Works

- [ ] Cursor is moving on your screen
- [ ] Status shows "Active" (green badge)
- [ ] Logs show "micro_move" entries
- [ ] When you move the mouse, logs show "user_active_pause"

### Troubleshooting Local Setup

| Problem | Solution |
|---------|----------|
| "Failed to fetch" from dashboard | Agent not running; check Terminal 1 |
| Cursor not moving | Check status is "Active"; agent may not have Windows hooks initialized |
| "Platform not implemented" error | You're on macOS/Linux; local cursor movement not supported (try cloud mode) |
| Port 8787 already in use | Kill previous agent: `Get-Process -Name apc-agent \| Stop-Process` |

---

## Cloud + Vercel Setup

Deploy both agent and dashboard to the cloud. Useful for:
- Remote desktop automation (keep RDP sessions alive)
- Centralized monitoring of multiple machines
- Access dashboard from anywhere

### Part 1: Deploy Agent to Cloud

#### Option A: Fly.io (Recommended - Free Tier)

1. **Install Fly.io CLI**:
   ```powershell
   choco install flyctl
   ```

2. **Login**:
   ```bash
   flyctl auth login
   ```

3. **Deploy from agent directory**:
   ```bash
   cd agent
   flyctl launch
   # Name: apc-agent
   # Region: pick closest to you
   # Keep other defaults
   ```

4. **Enable cloud mode** (edit generated `fly.toml`):
   ```toml
   [env]
   APC_CLOUD_MODE = "true"
   APC_ALLOW_INSECURE = "true"
   APC_BIND = "0.0.0.0:8787"
   ```

5. **Deploy**:
   ```bash
   flyctl deploy
   ```

6. **Get your URL**:
   ```bash
   flyctl status
   # Look for the app URL, e.g., https://apc-agent-xyz.fly.dev
   ```

#### Option B: Other Cloud Providers

See [DEPLOYMENT_CLOUD.md](DEPLOYMENT_CLOUD.md) for:
- Render.com (free tier, sleeps after 15 min)
- Railway.app ($5/month)
- DigitalOcean ($5/month)

### Part 2: Deploy Dashboard to Vercel

1. **Push code to GitHub**:
   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin https://github.com/YOUR_USERNAME/idea2026
   git push -u origin main
   ```

2. **Visit Vercel**: https://vercel.com/new

3. **Import repository**:
   - Select your `idea2026` repo
   - Root directory: `web`
   - Build command: `npm run build`
   - Output directory: `dist`

4. **Add environment variables**:
   ```
   VITE_AGENT_ORIGIN = https://apc-agent-xyz.fly.dev
   VITE_AGENT_TOKEN = (leave empty unless you set APC_PAIRING_TOKEN)
   ```
   Replace with your actual cloud agent URL from Step 1.6.

5. **Deploy** → Dashboard is live at `https://apc-dashboard.vercel.app`

6. **Test**: Open dashboard, click "Refresh", verify status appears

### Cost Breakdown

| Service | Free Tier | Paid |
|---------|-----------|------|
| **Fly.io** (Agent) | Yes (shared CPU) | $7/month |
| **Vercel** (Dashboard) | Yes (unlimited) | $20/month (optional) |
| **Total** | **$0** | $7-27/month |

**Recommendation**: Start with free tier; upgrade only if needed.

---

## Hybrid Setup (Recommended for WFH)

Deploy agent locally (moves your cursor), dashboard to Vercel (access from anywhere).

1. **Run agent locally** (follow "Run Locally" section)
2. **Deploy dashboard to Vercel** (follow Part 2 above)
3. **In Vercel environment variables**:
   ```
   VITE_AGENT_ORIGIN = http://YOUR_LOCAL_IP:8787
   ```
   Find `YOUR_LOCAL_IP`:
   ```powershell
   ipconfig  # Look for IPv4 address, e.g., 192.168.1.100
   ```

4. **Firewall**: Allow port 8787 inbound (or use ngrok tunnel)

**Benefits**:
- ✅ Cursor moves on your actual machine
- ✅ Dashboard accessible from mobile, other devices
- ✅ All control on local agent (no cloud data)
- ✅ Only dashboard in cloud (can be free Vercel)

---

## API Basics

All endpoints support optional Bearer token authentication:

```bash
curl -H "Authorization: Bearer YOUR_TOKEN" \
  http://127.0.0.1:8787/status
```

### Health Check (No Auth Required)
```bash
GET /health
# Response: {"status":"ok"}
```

### Get Status
```bash
GET /status
# Response: {"state":"active","mode":"meeting",...}
```

### Start Session
```bash
POST /session/start
Body: {
  "mode": "meeting",
  "randomness": 0.5,
  "idleThresholdSeconds": 2,
  "maxDurationMinutes": 240
}
```

### Stop Session
```bash
POST /session/stop
```

### View Logs
```bash
GET /logs
# Response: [{"ts":"...","action":"engine_start",...}]
```

### Live Status Stream (SSE)
```bash
GET /stream
# Receives server-sent events every 2 seconds
```

See [api.md](api.md) for full API reference.

---

## Next Steps

- **Scheduling**: Configure recurring sessions (API exists, UI not yet built)
- **Custom Modes**: Adjust movement speed, idle thresholds via config
- **Advanced Cloud**: Set up remote desktop API integration for true cloud cursor movement
- **Security**: Generate TLS certificates, set pairing tokens for production

See [architecture.md](architecture.md) for system design details.

---

## FAQ

### Can this get me in trouble at work?

This tool simulates *normal user activity* (mouse movements, preventing sleep). Most companies allow this explicitly—check your company's policy. The agent logs all actions for auditing.

### Does this work on macOS/Linux?

Not yet. Windows implementation is complete; macOS/Linux require platform-specific code (CGEvent, uinput). Contributions welcome!

### Can the cloud agent move my local cursor?

Not directly. Cloud agent is headless (no display output). To control your local mouse from the cloud, you'd need:
- An RDP/VNC connection (complex setup)
- Or local agent + remote dashboard (hybrid setup, recommended)

### Is my data safe?

Yes. Local setup stores everything on your machine. Cloud setup only logs *timing and actions*, no personal data. No telemetry or cloud sync.

### Can I use this on multiple machines?

Yes. Run local agent on each machine, point them all to the same cloud dashboard (or separate dashboards).

---

## Support

- **Issues**: Open a GitHub issue
- **Questions**: Check docs/ folder
- **Contributions**: PRs welcome

