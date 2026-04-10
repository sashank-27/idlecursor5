# 🎉 Activity Presence Controller - Cloud Agent Implementation Complete!

Welcome to your fully cloud-ready Activity Presence Controller! This project now supports both local cursor movement on Windows and headless cloud deployment.

---

## ⚡ Quick Start (Choose One)

### 🖥️ **I want to keep my Windows PC active (WFH)**
→ **5 minute setup** (Python version)

**Easiest Method - Automatic Setup:**
```batch
full-setup.bat
```
Then open `launcher.bat` to run the services.

**Manual Method:**
```batch
cd agent
setup-and-run.bat

# In another terminal:
cd web
npm-setup-and-run.bat
# Open http://localhost:5173 and click Start
```

**Next:** [Quick Start Windows Guide](QUICK_START_WINDOWS.md) | [Detailed Setup](docs/GETTING_STARTED.md#run-locally)

---

### ☁️ **I want to deploy to the cloud for free**
→ **20 minute setup**

```bash
# Deploy agent to Fly.io free tier
cd agent
flyctl launch
flyctl deploy

# Deploy dashboard to Vercel free tier
# Push to GitHub → vercel.com/new → Select repo → Deploy
```

**Next:** [Cloud Deployment Guide](docs/DEPLOYMENT_CLOUD.md)

---

### 🔄 **I want cursor movement + remote dashboard (Hybrid)**
→ **Recommended setup for teams**

Local cursor movement + Vercel dashboard = Full remote control

**Next:** [Hybrid Setup Guide](docs/GETTING_STARTED.md#hybrid-setup-recommended-for-wfh)

---

## 📚 Documentation

| Resource | Time | Purpose |
|----------|------|---------|
| [QUICK_REFERENCE.md](QUICK_REFERENCE.md) | 5 min | All features in one page |
| [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md) | 15 min | Detailed setup for all modes |
| [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md) | 20 min | Deploy to Fly.io, Render, Railway, DigitalOcean |
| [docs/DEPLOYMENT_VERCEL.md](docs/DEPLOYMENT_VERCEL.md) | 10 min | Deploy dashboard to Vercel |
| [docs/api.md](docs/api.md) | 10 min | REST API reference |
| [docs/architecture.md](docs/architecture.md) | 15 min | System design & security |
| [DOCUMENTATION_MAP.md](DOCUMENTATION_MAP.md) | 5 min | Navigation guide for all docs |

---

## ✨ What's New (This Implementation)

### Cloud Mode Support ☁️
```bash
APC_CLOUD_MODE=true   # Run on headless servers (no cursor movement)
APC_CLOUD_MODE=false  # Run on Windows (moves real cursor)
```

### Docker Ready 🐳
```dockerfile
# Deploy anywhere: Fly.io, Render, Railway, DigitalOcean
# Just push code → provider deploys automatically
docker build -t apc-agent .
docker run -e APC_CLOUD_MODE=true apc-agent
```

### Free Hosting 💰
```
Agent:     Fly.io free tier          → $0/month
Dashboard: Vercel free tier          → $0/month
Total:     Full cloud setup          → $0/month
```

### Bearer Token Auth 🔐
```bash
# Secure your agent with optional token
APC_PAIRING_TOKEN=your-secret-token
VITE_AGENT_TOKEN=your-secret-token  # In dashboard
```

---

## 🚀 Deployment Options

| Option | Cost | Setup Time | Best For |
|--------|------|-----------|----------|
| **Local Only** | $0 | 5 min | Single Windows PC |
| **Cloud Only** | $0-5/mo | 20 min | Remote servers |
| **Hybrid** | $0 | 15 min | Distributed teams |

---

## 🎯 What This Project Does

**Local Mode (Windows)**
- ✅ Moves your cursor automatically every 500ms
- ✅ Prevents system from sleeping
- ✅ Pauses when you move the mouse manually
- ✅ Keeps Teams/Slack/Zoom showing "active"
- ✅ Fully local - nothing leaves your PC

**Cloud Mode (Servers)**
- ✅ Runs on headless cloud servers
- ✅ Logs all activity for monitoring
- ✅ Same REST API as local mode
- ✅ Remote control via web dashboard
- ✅ No cursor output (server has no display)

---

## 💻 Technology Stack

```
Backend:   Go 1.21 (lightweight, single binary)
Frontend:  React 18 + Vite + TypeScript
Platform:  Windows (SendInput), macOS/Linux stubs
Hosting:   Fly.io, Render, Railway, DigitalOcean
Dashboard: Vercel (free)
```

---

## 🔐 Security & Privacy

- ✅ **Local-first**: All data stays on your device
- ✅ **No telemetry**: We don't know what you're doing
- ✅ **Open source**: Review the code yourself
- ✅ **Optional auth**: Bearer token authentication available
- ✅ **Transparent**: Visible when active, can't run silently

---

## 📊 Project Structure

```
idea2026/
├── 📄 README.md                      Overview
├── 📄 QUICK_REFERENCE.md             Cheat sheet
├── 📄 DOCUMENTATION_MAP.md            Navigation guide
├── 📄 IMPLEMENTATION_COMPLETE.md      What was just added
│
├── agent/                            Go backend
│   ├── Dockerfile                   Cloud deployment
│   ├── fly.toml.example             Fly.io config
│   └── internal/                    Business logic
│
├── web/                             React dashboard
│   └── src/                        UI components
│
└── docs/                            Guides
    ├── GETTING_STARTED.md          All setup modes
    ├── DEPLOYMENT_CLOUD.md         Cloud providers
    ├── DEPLOYMENT_VERCEL.md        Frontend hosting
    ├── api.md                      REST API
    └── architecture.md             System design
```

---

## 🧪 Verify Installation

After setup, test with:

```bash
# Check agent is running
curl http://127.0.0.1:8787/health
# Expected: {"status":"ok"}

# Check status
curl http://127.0.0.1:8787/status
# Expected: JSON with agent state

# View dashboard
open http://localhost:5173
```

---

## ❓ FAQ

**Q: Will my employer notice?**
A: This simulates normal user activity. Check your company's policy. The agent logs everything for auditing.

**Q: Can I use this on Mac/Linux?**
A: Local cursor movement is Windows-only (for now). Cloud mode works on any OS.

**Q: How much does it cost?**
A: Free! Use Fly.io free tier (agent) + Vercel free tier (dashboard). Optional paid tiers available ($5-7/mo).

**Q: Is it safe?**
A: Yes. It's just mouse movement and sleep prevention. You control what it does.

**Q: Can the cloud agent move my local cursor?**
A: Not directly. But you can run a local agent on your PC and control it from anywhere via Vercel dashboard (hybrid setup).

---

## 🚦 Getting Help

**Not sure where to start?**
→ Read [DOCUMENTATION_MAP.md](DOCUMENTATION_MAP.md) - It shows exactly what to read for your use case.

**Want to deploy to cloud?**
→ [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md) - Step by step for each provider.

**Having problems?**
→ [QUICK_REFERENCE.md#-troubleshooting](QUICK_REFERENCE.md#-troubleshooting) - Common issues & fixes.

**Want to understand the API?**
→ [docs/api.md](docs/api.md) - Full REST API reference.

---

## 🎓 Recommended Learning Path

### Day 1: Get It Working (30 min)
- [ ] Read [README.md](README.md) (2 min)
- [ ] Follow [Local Setup](docs/GETTING_STARTED.md#run-locally) (10 min)
- [ ] See cursor move on your screen (5 min)
- [ ] Check the logs in dashboard (5 min)

### Day 2: Deploy to Cloud (1 hour)
- [ ] Pick a provider: [Fly.io recommended](docs/DEPLOYMENT_CLOUD.md#option-1-flyio-recommended)
- [ ] Follow deployment steps (30 min)
- [ ] Deploy dashboard to Vercel (20 min)
- [ ] Access from your phone (10 min)

### Day 3: Understand the Details (1-2 hours)
- [ ] Read [architecture.md](docs/architecture.md)
- [ ] Review [api.md](docs/api.md)
- [ ] Explore code in `agent/internal/` and `web/src/`

---

## 📦 What You Get

- ✅ **Fully functional app** - Local cursor movement works right now
- ✅ **Cloud ready** - Deploy anywhere in 20 minutes
- ✅ **Well documented** - 5+ guides for every use case
- ✅ **Free hosting** - Fly.io + Vercel = $0/month
- ✅ **Secure** - Optional authentication, transparent logging
- ✅ **Open source** - Review & modify the code
- ✅ **Easy to deploy** - Docker + cloud provider = automated

---

## 🎯 Next Steps

1. **Choose your setup** from the "Quick Start" options above
2. **Follow the guide** for your chosen option
3. **Run the commands** - Takes 5-20 minutes
4. **See it work** - Cursor moves or dashboard loads
5. **Customize** as needed (optional)

---

## 🤝 Contributing

Found an issue? Have a better idea? Contributions welcome!

---

## 📜 License

MIT License - Use freely, modify, distribute. See [LICENSE](LICENSE) file.

---

## 🎉 You're All Set!

Everything is ready. Pick a quick start option above and get going!

**Questions?** Check [DOCUMENTATION_MAP.md](DOCUMENTATION_MAP.md) for the right guide.

**Status**: ✅ Agent builds | ✅ Dashboard builds | ✅ Docker ready | ✅ Documentation complete

