# Web Dashboard Deployment (Vercel)

The React/Vite dashboard can be deployed to **Vercel for free**. This guide covers deploying the frontend to control either a local or cloud agent.

## Prerequisites

- GitHub account (for auto-deployment)
- Vercel account (free at https://vercel.com)
- Your `apc-agent` backend running (local or cloud)

## Step 1: Prepare Code

### Option A: Fork/Push to GitHub

```bash
cd /path/to/idea2026
git init
git add .
git commit -m "Initial APC deployment"
git remote add origin https://github.com/YOUR_USERNAME/idea2026
git push -u origin main
```

### Option B: GitHub CLI

```bash
gh repo create idea2026 --public --source=. --remote=origin --push
```

## Step 2: Deploy to Vercel

### Via Web UI (Easiest)

1. Visit https://vercel.com/new
2. Click "Import Git Repository"
3. Select your `idea2026` repo
4. Configure import settings:
   - **Project name**: `apc-dashboard`
   - **Framework preset**: `Vite`
   - **Root directory**: `web`
   - **Build command**: `npm run build`
   - **Output directory**: `dist`

5. Click "Environment Variables" and add:
   ```
   VITE_AGENT_ORIGIN = https://apc-agent.fly.dev
   VITE_AGENT_TOKEN = (optional, only if cloud agent has pairing token)
   ```
   Replace `apc-agent.fly.dev` with your cloud agent URL.

6. Click "Deploy" → Wait for build to complete

7. Your dashboard is live at: `https://apc-dashboard.vercel.app`

### Via Vercel CLI

```bash
npm install -g vercel
cd web
vercel --prod
# Follow prompts; when asked for env vars:
# VITE_AGENT_ORIGIN: https://apc-agent.fly.dev
# VITE_AGENT_TOKEN: (leave blank if no token)
```

## Step 3: Configure Agent URL

### For Local Agent (WFH Use)

If your dashboard is on Vercel and agent is on your local machine:

1. **Your local agent** must be accessible to your browser
2. **Firewall**: Allow port 8787 inbound (or use ngrok tunnel)
3. In Vercel **Environment Variables**:
   ```
   VITE_AGENT_ORIGIN = http://YOUR_LOCAL_IP:8787
   ```
   Example: `http://192.168.1.100:8787`

4. **Redeploy**: Click "Deployments" → Latest → "Redeploy"

### For Cloud Agent (Remote Desktop)

If your agent is deployed to cloud (Fly.io, Render, Railway):

1. Get your agent URL from provider dashboard:
   - **Fly.io**: `https://apc-agent.fly.dev` (or your custom domain)
   - **Render**: `https://apc-agent.onrender.com`
   - **Railway**: URL from Railway dashboard
   - **DigitalOcean**: URL from App Platform dashboard

2. In Vercel **Environment Variables**:
   ```
   VITE_AGENT_ORIGIN = https://apc-agent.fly.dev
   VITE_AGENT_TOKEN = your-pairing-token (if set)
   ```

3. **Redeploy** the dashboard

## Step 4: Test Connection

1. Open your Vercel dashboard: `https://apc-dashboard.vercel.app`
2. Check browser console (F12) for errors
3. Click "Refresh" button
4. If connected, you'll see agent status
5. If error, check:
   - Agent is running
   - URL is correct
   - CORS allows your Vercel domain (cloud agents do by default)

## Using Custom Domain

### On Vercel

1. Click your project → Settings → Domains
2. Add your custom domain (e.g., `apc.yourdomain.com`)
3. Follow DNS setup instructions

### On Your DNS Provider

Add CNAME record:
```
apc  CNAME  cname.vercel.com
```

## Environment Variables Reference

| Variable | Default | Purpose |
|----------|---------|---------|
| `VITE_AGENT_ORIGIN` | `http://127.0.0.1:8787` | Agent backend URL (local or cloud) |
| `VITE_AGENT_TOKEN` | (empty) | Optional Bearer token for authentication |

## Troubleshooting

### "Failed to connect to agent"

- ✅ Agent is running and accessible at the URL
- ✅ CORS is enabled (cloud agents have this by default)
- ✅ Firewall allows HTTPS (port 443 from Vercel to your agent)

### "CORS error: not allowed"

- Your local agent may need to allow Vercel domain
- In agent, check CORS middleware allows `vercel.app` origins
- Or deploy agent behind a reverse proxy with CORS headers

### "Bearer token rejected"

- ✅ Token matches `APC_PAIRING_TOKEN` env var on agent
- ✅ Check token has no extra spaces

### Blank page on load

- Check browser DevTools → Network tab
- Verify `VITE_AGENT_ORIGIN` is correct
- Check agent is running: `curl https://<agent-url>/health`

## Advanced: Monorepo Setup

If you want the dashboard and agent in the same repository on Vercel:

1. **Create `vercel.json`** at project root:
   ```json
   {
     "buildCommand": "npm run build:all",
     "outputDirectory": "web/dist"
   }
   ```

2. **Update package.json** root:
   ```json
   {
     "scripts": {
       "build:all": "cd agent && go build ./cmd/apc-agent && cd ../web && npm run build"
     }
   }
   ```

3. Deploy to Vercel as normal

(Note: Vercel's Go support is limited; consider separate repositories for reliability)

## Cost

- **Vercel Free**: Unlimited deployments, domains, bandwidth
- **Vercel Pro**: $20/month (optional for custom analytics)

**Recommendation**: Use free tier; upgrade only if you need advanced features.

