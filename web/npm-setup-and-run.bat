@echo off
REM Activity Presence Controller - Dashboard Setup & Run Script (Batch)
REM Run this from the web directory: npm-setup-and-run.bat

setlocal enabledelayedexpansion

echo.
echo ================================================
echo Activity Presence Controller - Web Dashboard Setup
echo ================================================
echo.

REM Check if Node is installed
echo [*] Checking for Node.js installation...
node --version >nul 2>&1
if errorlevel 1 (
    echo.
    echo [X] Node.js is not installed or not in PATH
    echo.
    echo Please install Node.js from:
    echo https://nodejs.org/ (LTS version recommended)
    echo.
    echo After installation, restart this script
    echo.
    pause
    exit /b 1
)

REM Get Node version
for /f "tokens=*" %%i in ('node --version 2^>^&1') do set NODE_VERSION=%%i
echo [+] Found: %NODE_VERSION%

REM Check npm
npm --version >nul 2>&1
if errorlevel 1 (
    echo [X] npm is not installed
    pause
    exit /b 1
)

REM Get npm version
for /f "tokens=*" %%i in ('npm --version 2^>^&1') do set NPM_VERSION=%%i
echo [+] npm version: %NPM_VERSION%
echo.

REM Check if package.json exists
if not exist package.json (
    echo [X] package.json not found in current directory
    echo.
    pause
    exit /b 1
)

REM Install dependencies
echo [*] Installing npm dependencies...
echo This may take a few minutes...
echo.
npm install

if errorlevel 1 (
    echo.
    echo [X] Failed to install dependencies
    echo.
    pause
    exit /b 1
)

echo.
echo [+] All dependencies installed
echo.

REM Run the dashboard
echo.
echo ================================================
echo Starting Dashboard
echo ================================================
echo.
echo [*] Starting development server on http://localhost:5173
echo.
echo IMPORTANT: Make sure the agent is running!
echo In another terminal, navigate to the agent directory and run:
echo     cd ..\agent
echo     setup-and-run.bat
echo.
echo Press Ctrl+C to stop the dashboard
echo.

npm run dev

endlocal
pause
