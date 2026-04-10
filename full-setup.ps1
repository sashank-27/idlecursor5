# Activity Presence Controller - Full Setup Script
# This script sets up both agent and dashboard on a fresh Windows machine
# Run from PowerShell as: powershell -ExecutionPolicy Bypass -File full-setup.ps1

$ErrorActionPreference = "Stop"

function Write-Header {
    param([string]$Text)
    Write-Host "`n================================================" -ForegroundColor Cyan
    Write-Host $Text -ForegroundColor Cyan
    Write-Host "================================================`n" -ForegroundColor Cyan
}

function Write-Success {
    param([string]$Text)
    Write-Host "✓ $Text" -ForegroundColor Green
}

function Write-Error-Custom {
    param([string]$Text)
    Write-Host "✗ $Text" -ForegroundColor Red
}

function Write-Info {
    param([string]$Text)
    Write-Host "ℹ $Text" -ForegroundColor Yellow
}

Write-Header "Activity Presence Controller - Full Setup"

# Check Python
Write-Info "Checking Python..."
try {
    $pythonOutput = & python --version 2>&1
    Write-Success "Python found: $pythonOutput"
}
catch {
    Write-Error-Custom "Python not found. Install from https://www.python.org/downloads/"
    Write-Info "Make sure to check 'Add Python to PATH' during installation"
    exit 1
}

# Check Node
Write-Info "Checking Node.js..."
try {
    $nodeOutput = & node --version 2>&1
    $npmOutput = & npm --version 2>&1
    Write-Success "Node found: $nodeOutput, npm: $npmOutput"
}
catch {
    Write-Error-Custom "Node.js not found. Install from https://nodejs.org/"
    exit 1
}

# Setup Agent
Write-Header "Setting up Agent"

$agentPath = "agent"
if (-not (Test-Path $agentPath)) {
    Write-Error-Custom "agent directory not found"
    exit 1
}

Push-Location $agentPath

# Create venv
if (-not (Test-Path "venv")) {
    Write-Info "Creating Python virtual environment..."
    & python -m venv venv
    Write-Success "Virtual environment created"
}

# Activate venv
$activateScript = "venv\Scripts\Activate.ps1"
if (Test-Path $activateScript) {
    Write-Info "Activating virtual environment..."
    & $activateScript
    Write-Success "Virtual environment activated"
}

# Install Python dependencies
Write-Info "Installing Python dependencies..."
& pip install --upgrade pip
& pip install -r requirements.txt
Write-Success "Python dependencies installed"

Pop-Location

# Setup Dashboard
Write-Header "Setting up Dashboard"

$webPath = "web"
if (-not (Test-Path $webPath)) {
    Write-Error-Custom "web directory not found"
    exit 1
}

Push-Location $webPath

Write-Info "Installing npm dependencies..."
& npm install
Write-Success "npm dependencies installed"

Pop-Location

Write-Header "Setup Complete!"

Write-Info "To start the services, open TWO separate PowerShell windows:"
Write-Host "`n1. AGENT - In first window:"
Write-Host "   cd agent" -ForegroundColor Cyan
Write-Host "   .\venv\Scripts\Activate.ps1" -ForegroundColor Cyan
Write-Host "   `$env:APC_ALLOW_INSECURE='true'" -ForegroundColor Cyan
Write-Host "   `$env:APC_CLOUD_MODE='false'" -ForegroundColor Cyan
Write-Host "   python main.py" -ForegroundColor Cyan

Write-Host "`n2. DASHBOARD - In second window:"
Write-Host "   cd web" -ForegroundColor Cyan
Write-Host "   npm run dev" -ForegroundColor Cyan

Write-Host "`n3. Then open http://localhost:5173 in your browser`n" -ForegroundColor Cyan

Write-Info "Press any key to continue..."
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
