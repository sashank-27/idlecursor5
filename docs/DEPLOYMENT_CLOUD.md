# Cloud Deployment Guide

This guide explains how to deploy the Activity Presence Controller agent to free or low-cost cloud providers.

## Overview

The agent can run in two modes:

1. **Local Mode** (`APC_CLOUD_MODE=false`): Runs on your machine, moves your actual cursor via `SendInput` (Windows) or X11 (Linux)
2. **Cloud Mode** (`APC_CLOUD_MODE=true`): Runs on a remote server, logs activity but **does NOT move a cursor** (headless)

### Key Difference

- **Local Agent**: Keeps your local machine "active" for Teams/Slack/Zoom
- **Cloud Agent**: Useful for keeping remote desktop sessions alive (RDP/VNC), server health monitoring, or if you want to deploy a "presence controller" service that others access

## Deployment Options

### Option 1: Fly.io (Recommended - Free Tier Available)

**Pros**: Free tier with shared CPU, easy deployment, global deployment
**Cons**: Free tier has limited resources, no credits required

#### Steps

1. **Install Fly.io CLI**:
   ```powershell
   choco install flyctl
   # or download from https://fly.io/docs/getting-started/installing-flyctl/
   ```

2. **Login**:
   ```bash
   flyctl auth login
   ```

3. **Create app** (from agent directory):
   ```bash
   flyctl launch
   # When prompted:
   # - App name: apc-agent (or your choice)
   # - Region: pick closest to you
   # - Database: no
   # - Dockerfile: yes (should auto-detect)
   ```

4. **Set environment variables** (edit `fly.toml`):
   ```toml
   [env]
   APC_ALLOW_INSECURE = "true"
   APC_CLOUD_MODE = "true"
   APC_BIND = "0.0.0.0:8787"
   ```

5. **Deploy**:
   ```bash
   flyctl deploy
   ```

6. **Access**:
   ```bash
   flyctl apps list
   # Your app is at: https://<app-name>.fly.dev/health
   ```

---

### Option 2: Render.com (Free Tier - Sleeps after inactivity)

**Pros**: Free tier with GitHub auto-deploy
**Cons**: Free tier app spins down after 15 min of inactivity

#### Steps

1. **Push code to GitHub**:
   ```bash
   git init
   git add .
   git commit -m "Initial commit"
   git remote add origin https://github.com/YOUR_USERNAME/apc-agent
   git push -u origin main
   ```

2. **Visit [render.com](https://render.com)** and sign up with GitHub

3. **Create new Web Service**:
   - Select your repository
   - Name: `apc-agent`
   - Root directory: `agent`
   - Runtime: `Docker`
   - Plan: `Free`

4. **Set environment variables**:
   ```
   APC_ALLOW_INSECURE=true
   APC_CLOUD_MODE=true
   APC_BIND=0.0.0.0:8787
   ```

5. **Deploy** (auto-deploys on push to main)

6. **Access**:
   ```
   https://apc-agent.onrender.com/health
   ```

---

### Option 3: Railway.app ($5/month minimum)

**Pros**: Simple GitHub integration, good UI, good support
**Cons**: Requires payment (but free tier available)

#### Steps

1. **Push code to GitHub** (same as Render)

2. **Visit [railway.app](https://railway.app)** and sign up with GitHub

3. **Create new project**:
   - Deploy from GitHub repo
   - Select your `apc-agent` repo
   - Railway auto-detects Dockerfile

4. **Set variables** in Railway dashboard:
   ```
   APC_ALLOW_INSECURE=true
   APC_CLOUD_MODE=true
   APC_BIND=0.0.0.0:8787
   ```

5. **Deploy** (auto on push to main)

6. **Access**: Railway assigns a URL in the Networking section

---

### Option 4: DigitalOcean ($5/month App Platform)

**Pros**: Reliable, good performance, simple billing
**Cons**: Paid (cheapest option is $5/month)

#### Steps

1. **Create DigitalOcean account** and add payment method

2. **Visit Apps** → **Create App**:
   - Connect GitHub repo
   - Branch: `main`
   - Source: `agent` directory
   - Autodeploy: enabled

3. **Select `Dockerfile`** when prompted

4. **Set environment variables**:
   ```
   APC_ALLOW_INSECURE=true
   APC_CLOUD_MODE=true
   APC_BIND=0.0.0.0:8787
   ```

5. **Deploy**

6. **Access**: DigitalOcean assigns a domain (e.g., `apc-agent-xxxxx.ondigitalocean.app`)

---

## Accessing Your Deployed Agent

Once deployed, your agent is accessible at:
```
https://<app-url>/health
```

### From the dashboard (React frontend):

1. **Set `VITE_AGENT_ORIGIN`** environment variable when building:
   ```bash
   cd web
   VITE_AGENT_ORIGIN=https://<app-url> npm run build
   ```

2. **Deploy the `web/dist/` folder** to Vercel (free):
   ```bash
   npm run build
   # then push to GitHub and connect Vercel
   ```

3. **Access dashboard** at your Vercel URL and control the cloud agent

### Example Configuration

**Local Machine** (Windows):
```bash
set APC_BIND=127.0.0.1:8787
set APC_ALLOW_INSECURE=true
set APC_CLOUD_MODE=false
# Cursor will move on your actual machine
```

**Cloud Server**:
```bash
export APC_BIND=0.0.0.0:8787
export APC_ALLOW_INSECURE=true
export APC_CLOUD_MODE=true
# Cursor movement disabled; activity logged to /logs endpoint
```

---

## API Endpoints

All endpoints require `Authorization: Bearer <APC_PAIRING_TOKEN>` (if set):

| Endpoint | Method | Purpose |
|----------|--------|---------|
| `/health` | GET | Liveness check (no auth) |
| `/status` | GET | Current agent state |
| `/session/start` | POST | Start presence simulation |
| `/session/stop` | POST | Stop presence simulation |
| `/logs` | GET | View activity logs |
| `/stream` | GET | Server-sent events (live status updates) |

### Example Curl

```bash
curl -H "Authorization: Bearer mytoken" \
  https://apc-agent-xyz.fly.dev/status

# Start session
curl -X POST \
  -H "Authorization: Bearer mytoken" \
  https://apc-agent-xyz.fly.dev/session/start
```

---

## Monitoring & Logs

### Fly.io
```bash
flyctl logs
flyctl ssh console  # SSH into running app
```

### Render
Dashboard → Logs tab shows real-time logs

### Railway
Dashboard → Logs tab shows real-time logs

### DigitalOcean
Dashboard → View App → Logs tab

---

## Troubleshooting

### Agent not starting
- Check logs for errors
- Verify `APC_BIND` matches exposed port (should be `0.0.0.0:8787`)
- Ensure `APC_ALLOW_INSECURE=true` for dev testing

### Can't access `/health` endpoint
- Check app is running (check provider's dashboard)
- Verify HTTPS URL (some providers enforce HTTPS)
- Check firewall rules allow port 8787

### Dashboard can't reach agent
- Ensure `VITE_AGENT_ORIGIN` points to deployed cloud URL
- Check CORS headers (agent allows all origins in cloud mode)
- Verify agent is running: `curl https://<app-url>/health`

---

## Advanced: Custom Domain

Most providers let you configure a custom domain:

**Fly.io**:
```bash
flyctl certs create apc-agent.yourdom.com
# Then update DNS records
```

**Render/Railway**: See their dashboard → Custom Domain settings

---

## Cost Summary

| Provider | Free Tier | Paid Tier |
|----------|-----------|-----------|
| Fly.io | Yes (shared CPU) | $7-14/mo |
| Render | Yes (sleeps after 15min) | $7/mo |
| Railway | Limited | $5/mo minimum |
| DigitalOcean | No | $5/mo (App Platform) |

**Recommendation**: Start with **Fly.io free tier** to test, then upgrade if needed.

---

## Note on Cursor Movement in Cloud

Cloud deployments **do not move a cursor** because:
1. Cloud servers are headless (no display output)
2. Even with a desktop environment, the cursor would move on the *server's desktop*, not your local machine
3. To move *your local cursor* from a cloud service, you'd need reverse SSH tunneling or remote control APIs (advanced)

**For WFH presence control**, deploy the **local agent** on your machine and use the cloud agent only for:
- Remote desktop session keep-alive
- Monitoring/logging of activity
- Centralized control panel (dashboard on cloud, local agent on machine)

