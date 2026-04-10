@echo off
REM Activity Presence Controller - Full Automatic Setup
REM Run this script to set up everything

setlocal enabledelayedexpansion

echo.
echo ================================================
echo Activity Presence Controller - Full Setup
echo ================================================
echo.

REM Check Python
echo [*] Checking Python...
python --version >nul 2>&1
if errorlevel 1 (
    echo [X] Python not found. Install from https://www.python.org/downloads/
    echo     Make sure to check "Add Python to PATH"
    pause
    exit /b 1
)

for /f "tokens=*" %%i in ('python --version 2^>^&1') do echo [+] %%i

REM Check Node
echo [*] Checking Node.js...
node --version >nul 2>&1
if errorlevel 1 (
    echo [X] Node.js not found. Install from https://nodejs.org/
    pause
    exit /b 1
)

for /f "tokens=*" %%i in ('node --version 2^>^&1') do echo [+] Node %%i
for /f "tokens=*" %%i in ('npm --version 2^>^&1') do echo [+] npm %%i
echo.

REM Setup Agent
echo [*] Setting up Agent...
cd agent

if not exist requirements.txt (
    echo [X] requirements.txt not found in agent directory
    exit /b 1
)

echo     Installing Python dependencies...
pip install --upgrade pip >nul 2>&1
pip install -r requirements.txt >nul 2>&1

if errorlevel 1 (
    echo [X] Failed to install Python dependencies
    pause
    exit /b 1
)

echo [+] Agent dependencies installed
cd ..

REM Setup Dashboard
echo [*] Setting up Dashboard...
cd web

if not exist package.json (
    echo [X] package.json not found in web directory
    exit /b 1
)

echo     Installing npm dependencies...
call npm install >nul 2>&1

if errorlevel 1 (
    echo [X] Failed to install npm dependencies
    pause
    exit /b 1
)

echo [+] Dashboard dependencies installed
cd ..

echo.
echo ================================================
echo Setup Complete!
echo ================================================
echo.
echo To run the application, open TWO command prompts:
echo.
echo 1. FIRST PROMPT - Agent:
echo    cd agent
echo    setup-and-run.bat
echo.
echo 2. SECOND PROMPT - Dashboard:
echo    cd web
echo    npm-setup-and-run.bat
echo.
echo Then open http://localhost:5173 in your browser
echo.
pause
