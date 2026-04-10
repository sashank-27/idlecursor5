# Web Dashboard Setup Scripts

This directory contains the React-based web dashboard for the Activity Presence Controller with automatic setup scripts for Windows.

## Quick Start

### Easiest Method (Automatic)
From this directory, double-click or run:
```batch
npm-setup-and-run.bat
```

This will:
1. Check if Node.js is installed
2. Install all npm dependencies
3. Start the development server on `http://localhost:5173`

### What's Running
- **Dashboard**: React PWA (Progressive Web App)
- **Port**: 5173 (development server)
- **Frontend**: Vite dev server with hot module replacement

## Files

### Scripts
- **npm-setup-and-run.bat** - Batch script for automatic setup and run

### React Files
- **src/App.tsx** - Main React component
- **src/api.ts** - API client for communicating with agent
- **src/main.tsx** - Entry point
- **src/style.css** - Styles

### Configuration
- **package.json** - npm dependencies
- **vite.config.ts** - Vite build configuration
- **tsconfig.json** - TypeScript configuration

## Prerequisites

### Node.js and npm
- Download from https://nodejs.org/ (LTS version recommended)
- After installing, verify: `node --version` and `npm --version`

## Environment Variables

The dashboard reads the agent URL from:
- Default: `http://localhost:8787` (local development)

To connect to a different agent, modify `src/api.ts`:
```typescript
const API_BASE_URL = "http://your-agent-url:8787";
```

## Available Scripts

### `npm install`
Install all project dependencies (done by setup script)

### `npm run dev`
Start the development server (done by setup script)
- Open http://localhost:5173 in browser
- Hot module replacement enabled (code changes auto-reload)

### `npm run build`
Build for production
- Outputs to `dist/` directory
- Ready for deployment to Vercel, Netlify, etc.

### `npm run preview`
Preview the production build locally

## Troubleshooting

### Node.js Not Found
- Install from https://nodejs.org/
- **Make sure to check "Add to PATH"**
- Close all terminals and try again

### Port Already in Use
```batch
netstat -ano | findstr :5173
taskkill /PID <pid> /F
```

### Connection to Agent Failed
- Make sure agent is running on `http://127.0.0.1:8787`
- Check Windows Firewall isn't blocking port 8787
- Look for errors in browser console (F12)

### npm install Fails
Try:
```batch
npm cache clean --force
npm install
```

### Hot reload not working
- Hard refresh: Ctrl+Shift+R
- Close browser and reopen http://localhost:5173

## Project Structure

```
src/
├── App.tsx           - Main component
├── api.ts           - API client
├── main.tsx         - Entry point
├── style.css        - Global styles
└── ... other components
public/
├── index.html       - HTML template
└── manifest.webmanifest - PWA manifest
```

## Building for Production

```batch
npm run build
```

This creates an optimized production build in the `dist/` folder.

To deploy to Vercel:
1. Push to GitHub
2. Go to https://vercel.com/new
3. Select this repository
4. Click "Deploy"

## Features

- Real-time status updates via SSE (Server-Sent Events)
- Start/stop activity sessions
- View action logs
- Policy lock/unlock
- Responsive design
- Offline-capable (PWA)

## API Contract

The dashboard expects these endpoints from the agent:

- `GET /health` - Health check
- `GET /status` - Current state
- `POST /session/start` - Start session
- `POST /session/stop` - Stop session
- `POST /policy/lock` - Lock/unlock
- `GET /logs` - Get logs
- `GET /stream` - Server-sent events

See `src/api.ts` for implementation details.
