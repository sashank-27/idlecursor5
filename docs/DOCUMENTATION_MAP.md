# Project Structure & Documentation Map

## 📁 Project Layout

```
idea2026/
├── 📄 README.md                          ← Start here! Overview & quick start
├── 📄 QUICK_REFERENCE.md                 ← Cheat sheet for common tasks
├── 📄 IMPLEMENTATION_COMPLETE.md          ← What was just implemented
├── 📄 CLOUD_IMPLEMENTATION.md             ← Technical details of cloud support
├── LICENSE                               ← MIT License
│
├── agent/                                ← Go backend (activity agent)
│   ├── Dockerfile                        ← Container build (NEW)
│   ├── .dockerignore                     ← Build optimization (NEW)
│   ├── fly.toml.example                  ← Fly.io config example (NEW)
│   ├── go.mod                            ← Go dependencies
│   ├── go.sum                            ← Dependency checksums
│   │
│   ├── cmd/
│   │   └── apc-agent/
│   │       └── main.go                   ← Entry point (MODIFIED)
│   │
│   └── internal/
│       ├── api/
│       │   └── server.go                 ← HTTP API & config (MODIFIED)
│       ├── behavior/
│       │   └── engine.go                 ← Activity simulation logic
│       └── system/
│           ├── platform.go               ← Platform interface
│           ├── platform_windows.go       ← Windows implementation (SendInput, etc)
│           └── platform_stub.go          ← Placeholder for other OS
│
├── web/                                  ← React/Vite frontend
│   ├── index.html                        ← Entry point
│   ├── manifest.webmanifest              ← PWA manifest
│   ├── package.json                      ← npm dependencies
│   ├── tsconfig.json                     ← TypeScript config
│   ├── vite.config.ts                    ← Vite config
│   │
│   └── src/
│       ├── main.tsx                      ← React entry
│       ├── App.tsx                       ← Main component
│       ├── api.ts                        ← Backend client (MODIFIED)
│       ├── style.css                     ← Dark theme styles
│       ├── vite-env.d.ts                 ← Env types (MODIFIED)
│       ├── react-shim.d.ts               ← React type stubs
│       └── dist/                         ← Build output (ready for CDN)
│
└── docs/                                 ← Documentation
    ├── api.md                            ← REST API reference
    ├── architecture.md                   ← System design
    ├── GETTING_STARTED.md                ← Comprehensive quick start (NEW)
    ├── DEPLOYMENT_CLOUD.md               ← Cloud provider guides (NEW)
    └── DEPLOYMENT_VERCEL.md              ← Vercel deployment (UPDATED)
```

---

## 📚 Documentation Guide

### For Different User Types

**👤 First Time User**
1. Read: [README.md](README.md) (2 min)
2. Choose setup: [QUICK_REFERENCE.md](QUICK_REFERENCE.md) (5 min)
3. Follow guide: [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md) (15 min)

**🚀 Want to Deploy to Cloud**
1. Read: [QUICK_REFERENCE.md](QUICK_REFERENCE.md) - "Deployment Providers" section
2. Choose provider: [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md)
3. Deploy frontend: [docs/DEPLOYMENT_VERCEL.md](docs/DEPLOYMENT_VERCEL.md)

**👨‍💻 Developer/Hacker**
1. Architecture: [docs/architecture.md](docs/architecture.md)
2. API Reference: [docs/api.md](docs/api.md)
3. Code: Browse `agent/internal/` and `web/src/`

**🏢 Enterprise/IT Admin**
1. Security Model: [docs/architecture.md](docs/architecture.md) - "Security & Privacy"
2. Deployment: [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md)
3. Configuration: [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md) - "API Basics"

**📖 Implementation Details**
- What was added: [IMPLEMENTATION_COMPLETE.md](IMPLEMENTATION_COMPLETE.md)
- Technical details: [CLOUD_IMPLEMENTATION.md](CLOUD_IMPLEMENTATION.md)

---

## 🔍 Find What You Need

### By Task

| Task | File |
|------|------|
| **Run locally on Windows** | [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md#run-locally) |
| **Deploy to cloud (free)** | [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md#cloud--vercel-setup) |
| **Deploy agent to Fly.io** | [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md#option-1-flyio-recommended) |
| **Deploy frontend to Vercel** | [docs/DEPLOYMENT_VERCEL.md](docs/DEPLOYMENT_VERCEL.md) |
| **Understand the API** | [docs/api.md](docs/api.md) |
| **Configure environment variables** | [QUICK_REFERENCE.md](QUICK_REFERENCE.md#-environment-variables) |
| **Troubleshoot connection issues** | [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md#troubleshooting-local-setup) or [QUICK_REFERENCE.md](QUICK_REFERENCE.md#-troubleshooting) |
| **See system architecture** | [docs/architecture.md](docs/architecture.md) |
| **Understand what's new** | [IMPLEMENTATION_COMPLETE.md](IMPLEMENTATION_COMPLETE.md) |

### By Question

**Q: "How do I get started?"**
→ [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md)

**Q: "What's the difference between local and cloud mode?"**
→ [QUICK_REFERENCE.md](QUICK_REFERENCE.md#-what-s-new) or [CLOUD_IMPLEMENTATION.md](CLOUD_IMPLEMENTATION.md#deployment-modes)

**Q: "How much does it cost?"**
→ [QUICK_REFERENCE.md](QUICK_REFERENCE.md) or [IMPLEMENTATION_COMPLETE.md](IMPLEMENTATION_COMPLETE.md#-cost-summary)

**Q: "Can I use this on macOS/Linux?"**
→ [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md#faq)

**Q: "What's the API?"**
→ [docs/api.md](docs/api.md)

**Q: "Is it secure?"**
→ [docs/architecture.md](docs/architecture.md#security--privacy-model)

**Q: "What was just added?"**
→ [IMPLEMENTATION_COMPLETE.md](IMPLEMENTATION_COMPLETE.md)

---

## 🎯 Recommended Reading Order

### For Local Setup (Windows WFH Use)
1. [README.md](README.md) - 2 min
2. [docs/GETTING_STARTED.md#run-locally](docs/GETTING_STARTED.md#run-locally) - 10 min
3. [QUICK_REFERENCE.md](QUICK_REFERENCE.md) - 5 min (bookmark for later)

### For Cloud + Vercel (Remote Access)
1. [README.md](README.md) - 2 min
2. [docs/GETTING_STARTED.md#cloud--vercel-setup](docs/GETTING_STARTED.md#cloud--vercel-setup) - 15 min
3. [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md) - 10-20 min (depends on provider)
4. [docs/DEPLOYMENT_VERCEL.md](docs/DEPLOYMENT_VERCEL.md) - 10 min

### For Developers
1. [docs/architecture.md](docs/architecture.md) - 15 min
2. [docs/api.md](docs/api.md) - 10 min
3. Code exploration: `agent/internal/` and `web/src/`
4. [CLOUD_IMPLEMENTATION.md](CLOUD_IMPLEMENTATION.md) - Technical reference

---

## 📋 File Reference

### New Files (This Implementation)
```
✨ CLOUD_IMPLEMENTATION.md           Technical details of cloud support
✨ IMPLEMENTATION_COMPLETE.md        Summary of changes
✨ QUICK_REFERENCE.md               Quick reference guide
✨ agent/Dockerfile                 Docker container build
✨ agent/.dockerignore               Build optimization
✨ agent/fly.toml.example            Fly.io config template
✨ docs/GETTING_STARTED.md           Comprehensive quick start
✨ docs/DEPLOYMENT_CLOUD.md          Cloud provider deployment guides
```

### Modified Files
```
⚡ README.md                         Added cloud deployment info
⚡ agent/cmd/apc-agent/main.go       Added APC_CLOUD_MODE support
⚡ agent/internal/api/server.go      Added CloudMode config
⚡ web/src/api.ts                    Added Bearer token support
⚡ web/src/vite-env.d.ts             Added VITE_AGENT_TOKEN type
```

### Existing Files (Unchanged)
```
✓ agent/internal/behavior/engine.go  Already handles nil platform
✓ agent/internal/system/platform.go  Interface unchanged
✓ agent/internal/system/platform_windows.go  Works as-is
✓ agent/internal/system/platform_stub.go     Works as-is
✓ web/src/App.tsx                    No changes needed
✓ web/src/style.css                  No changes needed
✓ docs/api.md                        Still valid
✓ docs/architecture.md               Still valid
```

---

## 🚀 Quick Start Commands

### Local Mode (Windows)
```powershell
# Terminal 1: Start agent
cd agent
$env:APC_CLOUD_MODE="false"
$env:APC_ALLOW_INSECURE="true"
go run ./cmd/apc-agent/main.go

# Terminal 2: Start dashboard
cd web
npm install
npm run dev
# → Open http://localhost:5173
```

### Cloud Mode (Fly.io)
```bash
cd agent
flyctl launch
# Edit fly.toml: APC_CLOUD_MODE=true
flyctl deploy

# Deploy dashboard to Vercel pointing to your Fly.io URL
```

### Verify Installation
```bash
# Test agent health
curl http://127.0.0.1:8787/health
# Expected: {"status":"ok"}

# Test API
curl http://127.0.0.1:8787/status
# Expected: JSON with agent status
```

---

## 📞 Support

| Issue | Resource |
|-------|----------|
| Getting started | [docs/GETTING_STARTED.md](docs/GETTING_STARTED.md) |
| Local setup problems | [docs/GETTING_STARTED.md#troubleshooting-local-setup](docs/GETTING_STARTED.md#troubleshooting-local-setup) |
| Cloud deployment | [docs/DEPLOYMENT_CLOUD.md](docs/DEPLOYMENT_CLOUD.md) |
| Vercel deployment | [docs/DEPLOYMENT_VERCEL.md](docs/DEPLOYMENT_VERCEL.md) |
| API usage | [docs/api.md](docs/api.md) |
| Architecture questions | [docs/architecture.md](docs/architecture.md) |
| General issues | [QUICK_REFERENCE.md#-troubleshooting](QUICK_REFERENCE.md#-troubleshooting) |

---

## ✨ Project Status

- ✅ Local cursor movement (Windows)
- ✅ Cloud headless mode
- ✅ Docker containerization
- ✅ Vercel frontend deployment
- ✅ Free tier hosting (Fly.io + Vercel)
- ✅ Comprehensive documentation
- 🟡 macOS/Linux (stubbed, not implemented)
- 🟡 TLS auto-generation (configured, not auto)
- ⏳ Scheduling UI (API exists, UI not built)
- ⏳ Remote desktop integration (advanced feature)

---

## 🎓 Learning Path

### Day 1: Get It Running
- [ ] Read [README.md](README.md)
- [ ] Run [local setup](docs/GETTING_STARTED.md#run-locally)
- [ ] See cursor move on your Windows machine
- [ ] Understand what it does

### Day 2: Deploy to Cloud
- [ ] Choose a provider ([Fly.io recommended](docs/DEPLOYMENT_CLOUD.md#option-1-flyio-recommended))
- [ ] Deploy agent to cloud
- [ ] Deploy dashboard to Vercel
- [ ] Access from phone/other device

### Day 3: Understand the Details
- [ ] Read [API reference](docs/api.md)
- [ ] Review [architecture](docs/architecture.md)
- [ ] Explore source code
- [ ] Consider customization options

### Beyond: Advanced Customization
- [ ] Set pairing tokens for security
- [ ] Configure TLS certificates
- [ ] Implement scheduling (API exists)
- [ ] Contribute macOS/Linux support

---

**🎯 Next Step:** Pick your use case from [QUICK_REFERENCE.md](QUICK_REFERENCE.md#-recommended-setups) and follow the corresponding guide!

