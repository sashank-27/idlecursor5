# Quick Reference - Cloud Agent with Cursor Movement

## 🚀 What's New

Your Activity Presence Controller now supports **cloud deployment** while maintaining full cursor movement on local machines. Choose your deployment:

### **Local Mode** (Windows: Cursor Moves)
```powershell
# Terminal 1: Start Agent
cd agent
$env:APC_CLOUD_MODE="false"
$env:APC_ALLOW_INSECURE="true"
go run ./cmd/apc-agent/main.go

# Terminal 2: Start Dashboard
cd web
npm install
npm run dev
# → Open http://localhost:5173
```

### **Cloud Mode** (Headless: No Cursor)
```bash
# Deploy to Fly.io (free):
cd agent
flyctl launch
# Edit fly.toml: APC_CLOUD_MODE=true
flyctl deploy

# Deploy dashboard to Vercel (free):
# Push to GitHub → Visit vercel.com/new → Select repo → Deploy
```

### **Hybrid** (Local Cursor + Remote Dashboard)
```bash
# Agent on your PC (local mode)
# Dashboard on Vercel (VITE_AGENT_ORIGIN=http://YOUR_IP:8787)
# → Control from mobile/anywhere
```

---

## 📋 Complete Feature List

| Feature | Local | Cloud | Notes |
|---------|-------|-------|-------|
| **Cursor Movement** | ✅ | ❌ | Cloud servers are headless |
| **Sleep Prevention** | ✅ | ✅ | Works on any OS |
| **Activity Logging** | ✅ | ✅ | All actions recorded |
| **User Pause Detection** | ✅ | ❌ | Requires cursor API |
| **Remote Dashboard** | ✅ | ✅ | Can be anywhere |
| **Cost** | Free | $0-5/mo | Local free, cloud $0-5 |

---

## 🏗️ Architecture

### Local Setup
```
Your PC (Windows)
  ├─ apc-agent (127.0.0.1:8787)
  │   └─ Moves your cursor via SendInput
  └─ Browser (localhost:5173)
      └─ Controls agent via HTTP
```

### Cloud Setup
```
Your PC (anywhere)
  ├─ Browser (vercel.app)
  │   └─ Controls agent via HTTPS
  └─ Cloud Server
      └─ apc-agent (0.0.0.0:8787)
          └─ No cursor (headless)
```

---

## 🔧 Environment Variables

### For Local Agent
```bash
APC_BIND=127.0.0.1:8787         # Only local connections
APC_CLOUD_MODE=false             # Move cursor (Windows)
APC_ALLOW_INSECURE=true          # Allow HTTP (dev only)
APC_PAIRING_TOKEN=(optional)     # Bearer token auth
```

### For Cloud Agent (Dockerfile)
```bash
APC_BIND=0.0.0.0:8787            # Accept from anywhere
APC_CLOUD_MODE=true              # Disable cursor (headless)
APC_ALLOW_INSECURE=true          # Allow HTTP (dev)
APC_PAIRING_TOKEN=(optional)     # Bearer token auth
```

### For Dashboard (Vercel)
```bash
VITE_AGENT_ORIGIN=http://127.0.0.1:8787     # Local
VITE_AGENT_ORIGIN=https://apc-agent.fly.dev # Cloud
VITE_AGENT_TOKEN=your-token                 # If agent requires auth
```

---

## 📦 Deployment Providers

### Fly.io (Recommended)
- **Cost**: Free tier + paid starting $7/month
- **Setup**: 3 commands (`flyctl launch` → edit `fly.toml` → `flyctl deploy`)
- **Speed**: Global CDN, fast startup

### Render.com
- **Cost**: Free (sleeps after 15 min) + $7/month paid
- **Setup**: GitHub auto-deploy, simple UI
- **Downside**: Free tier pauses frequently

### Railway
- **Cost**: $5/month minimum
- **Setup**: GitHub integration, clean dashboard
- **Upside**: Always active (even free)

### DigitalOcean
- **Cost**: $5/month
- **Setup**: App Platform, simple interface
- **Upside**: Reliable, good documentation

### Vercel (Dashboard Only)
- **Cost**: Free tier (unlimited)
- **Setup**: Push to GitHub → connect to Vercel → Done
- **Features**: Auto-deploy on push, custom domains

---

## 🎯 Recommended Setups

### For WFH (Teams/Slack Activity)
```
✅ Local agent (moves real cursor)
✅ Dashboard on Vercel (optional, remote access)
💰 Cost: $0 (free tier)
```

### For Remote Desktop Sessions
```
✅ Cloud agent (keep server awake)
✅ Monitor logs & activity from cloud dashboard
💰 Cost: $0-5/month
```

### For Multiple Machines
```
✅ Local agent on each machine
✅ Central dashboard on Vercel
✅ Control all from one place
💰 Cost: $0 (all free tier)
```

---

## 🧪 Testing

### Local Mode
1. Start agent: `APC_CLOUD_MODE=false` → cursor should move
2. Start dashboard: `npm run dev`
3. Click "Start" → Watch cursor jump around
4. Move your mouse → Agent pauses for 15s
5. Click "Stop" → Cursor stops moving

### Cloud Mode
1. Deploy agent: `APC_CLOUD_MODE=true`
2. Dashboard connects to cloud URL
3. Click "Start" → Logs show activity (no cursor output)
4. Check `/logs` endpoint: see "micro_move" events (even though no cursor on server)

---

## 📚 Documentation Files

| File | Purpose |
|------|---------|
| [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md) | Quick start for all modes |
| [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md) | Deploy to Fly.io, Render, Railway, DigitalOcean |
| [docs/DEPLOYMENT_VERCEL.md](docs/DEPLOYMENT_VERCEL.md) | Deploy dashboard to Vercel |
| [docs/architecture.md](docs/architecture.md) | System design & API |
| [docs/api.md](docs/api.md) | REST API reference |
| [CLOUD_IMPLEMENTATION.md](CLOUD_IMPLEMENTATION.md) | What was added in this update |

---

## 🐛 Troubleshooting

### "Failed to fetch" from Dashboard
→ Agent not running or wrong URL in `VITE_AGENT_ORIGIN`

### Cursor Not Moving
→ Check `APC_CLOUD_MODE=false` for local mode

### Port 8787 Already in Use
→ Kill previous agent: `Get-Process -Name apc-agent | Stop-Process`

### Can't Access Cloud Agent
→ Check deployment logs on Fly.io/Render/Railway dashboard

### Bearer Token Rejected
→ Verify `VITE_AGENT_TOKEN` matches `APC_PAIRING_TOKEN` on agent

---

## 🚢 Deployment Checklist

### Local Deployment
- [ ] Go 1.21+ installed
- [ ] Run agent with `APC_CLOUD_MODE=false`
- [ ] Dashboard connects to `http://127.0.0.1:8787`
- [ ] Cursor moves on screen
- [ ] Status shows "Active"

### Cloud Deployment
- [ ] GitHub account set up
- [ ] Code pushed to GitHub
- [ ] Fly.io (or provider) account created
- [ ] `flyctl launch` completed
- [ ] `APC_CLOUD_MODE=true` in config
- [ ] `flyctl deploy` successful
- [ ] Dashboard on Vercel pointing to cloud URL
- [ ] `/health` endpoint responds with `{"status":"ok"}`

---

## 💡 Pro Tips

1. **Use Hybrid Setup**: Local cursor + Vercel dashboard = best UX
2. **Set Pairing Token**: `APC_PAIRING_TOKEN=your-secret` for security
3. **Monitor Cloud Logs**: `flyctl logs` to debug issues
4. **Custom Domain**: Vercel allows free custom domain for dashboard
5. **Test Locally First**: Always test local mode before deploying to cloud

---

## 📞 Need Help?

- Check [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md) for detailed walkthroughs
- See [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md) for provider-specific steps
- Review [docs/api.md](docs/api.md) for API details
- Check provider dashboards for deployment logs (Fly.io, Vercel, etc.)

---

## ✨ What's Next?

- Implement scheduling UI (API exists, UI not yet built)
- Add macOS/Linux platform implementations
- TLS certificate auto-generation
- Remote desktop integration (advanced)
- Multi-agent control panel

**Start small, scale up:**
1. ✅ Try local mode on Windows
2. ✅ Deploy dashboard to Vercel
3. ✅ Add cloud agent to Fly.io
4. → Build multi-machine setup

