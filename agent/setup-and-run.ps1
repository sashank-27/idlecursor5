# Activity Presence Controller - Python Agent Setup & Run Script
# This script sets up all dependencies and runs the agent
# Run from PowerShell as: powershell -ExecutionPolicy Bypass -File setup-and-run.ps1

param(
    [bool]$UseVirtualEnv = $true,
    [string]$PythonVersion = "3.9",
    [string]$Port = "8787"
)

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

Write-Header "Activity Presence Controller - Setup & Run"

# Check if Python is installed
Write-Info "Checking if Python is installed..."
$pythonPath = $null
$pythonVersion = $null

try {
    $pythonOutput = & python --version 2>&1
    if ($LASTEXITCODE -eq 0) {
        $pythonPath = (Get-Command python).Source
        Write-Success "Python found: $pythonPath"
        Write-Success "Version: $pythonOutput"
    }
}
catch {
    $pythonPath = $null
}

if (-not $pythonPath) {
    Write-Error-Custom "Python is not installed or not in PATH"
    Write-Info "Please install Python 3.9 or later from https://www.python.org/downloads/"
    Write-Info "Make sure to check 'Add Python to PATH' during installation"
    exit 1
}

# Check Python version
$versionMatch = [regex]::Match((& python --version 2>&1), "(\d+\.\d+)")
if ($versionMatch.Success) {
    $installedVersion = [version]$versionMatch.Groups[1].Value
    $requiredVersion = [version]$PythonVersion
    
    if ($installedVersion -lt $requiredVersion) {
        Write-Error-Custom "Python version $installedVersion is too old. Need $requiredVersion or later."
        exit 1
    }
}

Write-Success "Python version check passed"

# Create virtual environment if requested
$venvPath = Join-Path (Get-Location) "venv"
if ($UseVirtualEnv -and -not (Test-Path $venvPath)) {
    Write-Info "Creating virtual environment..."
    & python -m venv venv
    Write-Success "Virtual environment created"
}

# Activate virtual environment if it exists
if ($UseVirtualEnv -and (Test-Path $venvPath)) {
    Write-Info "Activating virtual environment..."
    $activateScript = Join-Path $venvPath "Scripts\Activate.ps1"
    if (Test-Path $activateScript) {
        & $activateScript
        Write-Success "Virtual environment activated"
    }
}

# Check if requirements.txt exists
$requirementsPath = "requirements.txt"
if (-not (Test-Path $requirementsPath)) {
    Write-Error-Custom "requirements.txt not found in current directory"
    exit 1
}

# Install pip packages
Write-Header "Installing Python Dependencies"
Write-Info "Installing packages from requirements.txt..."

& pip install --upgrade pip
if ($LASTEXITCODE -ne 0) {
    Write-Error-Custom "Failed to upgrade pip"
    exit 1
}

& pip install -r requirements.txt
if ($LASTEXITCODE -ne 0) {
    Write-Error-Custom "Failed to install requirements"
    exit 1
}

Write-Success "All dependencies installed"

# Verify imports work
Write-Header "Verifying Installation"
Write-Info "Checking if all modules can be imported..."

$testScript = @"
try:
    import flask
    import flask_cors
    print("✓ Flask packages OK")
except ImportError as e:
    print(f"✗ Import error: {e}")
    exit(1)
"@

& python -c $testScript
if ($LASTEXITCODE -ne 0) {
    Write-Error-Custom "Module verification failed"
    exit 1
}

Write-Success "All modules verified"

# Set environment variables (optional - agent defaults to insecure for dev)
Write-Header "Starting Agent"
Write-Info "Setting environment variables..."

$env:APC_CLOUD_MODE = "false"
$env:APC_BIND = "127.0.0.1:$Port"

Write-Success "Environment variables set:"
Write-Host "  APC_CLOUD_MODE = $env:APC_CLOUD_MODE"
Write-Host "  APC_BIND = $env:APC_BIND"
Write-Host "  (APC_ALLOW_INSECURE defaults to true for development)"

# Run the agent
Write-Header "Running Agent"
Write-Info "Starting Python agent on http://127.0.0.1:$Port"
Write-Info "Press Ctrl+C to stop the agent"
Write-Info "`nOpen another PowerShell window and run:"
Write-Host "`n  cd web" -ForegroundColor Cyan
Write-Host "  npm install" -ForegroundColor Cyan
Write-Host "  npm run dev" -ForegroundColor Cyan
Write-Host "`nThen open http://localhost:5173 in your browser`n" -ForegroundColor Cyan

& python main.py

if ($LASTEXITCODE -ne 0) {
    Write-Error-Custom "Agent exited with error code $LASTEXITCODE"
    exit $LASTEXITCODE
}
