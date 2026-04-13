@echo off
REM Activity Presence Controller - Python Agent Setup & Run Script (Batch)
REM Run this from the agent directory: setup-and-run.bat

setlocal enabledelayedexpansion

echo.
echo ================================================
echo Activity Presence Controller - Setup ^& Run
echo ================================================
echo.

REM Check if Python is installed
echo [*] Checking for Python installation...
python --version >nul 2>&1
if errorlevel 1 (
    echo.
    echo [X] Python is not installed or not in PATH
    echo.
    echo Please install Python 3.9 or later from:
    echo https://www.python.org/downloads/
    echo.
    echo IMPORTANT: During installation, check the box that says
    echo "Add Python to PATH"
    echo.
    pause
    exit /b 1
)

REM Get Python version
for /f "tokens=*" %%i in ('python --version 2^>^&1') do set PYTHON_VERSION=%%i
echo [+] Found: %PYTHON_VERSION%
echo.

REM Check if requirements.txt exists
if not exist requirements.txt (
    echo [X] requirements.txt not found in current directory
    echo.
    pause
    exit /b 1
)

REM Install pip packages
echo [*] Installing Python dependencies...
echo.
pip install --upgrade pip
if errorlevel 1 (
    echo [X] Failed to upgrade pip
    pause
    exit /b 1
)

pip install -r requirements.txt
if errorlevel 1 (
    echo [X] Failed to install requirements
    pause
    exit /b 1
)

echo.
echo [+] All dependencies installed
echo.

REM Set environment variables (optional - agent defaults to insecure for dev)
echo [*] Setting environment variables...
set APC_CLOUD_MODE=false
set APC_BIND=127.0.0.1:8787

echo [+] Environment configured:
echo     APC_CLOUD_MODE = %APC_CLOUD_MODE%
echo     APC_BIND = %APC_BIND%
echo     (APC_ALLOW_INSECURE defaults to true for development)
echo.

REM Run the agent
echo.
echo ================================================
echo Starting Agent
echo ================================================
echo.
echo [*] Starting Python agent on http://127.0.0.1:8787
echo.
echo IMPORTANT: Open another command prompt or PowerShell and run:
echo     cd web
echo     npm install
echo     npm run dev
echo.
echo Then open http://localhost:5173 in your browser
echo.
echo Press Ctrl+C to stop the agent
echo.

python main.py

endlocal
pause
