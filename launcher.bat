@echo off
REM Activity Presence Controller - Quick Start Launcher
REM This script provides an interactive menu to set up and run the application

setlocal enabledelayedexpansion

:menu
cls
echo.
echo ================================================
echo Activity Presence Controller - Launcher
echo ================================================
echo.
echo 1. Full Setup (install all dependencies)
echo 2. Run Agent Only
echo 3. Run Dashboard Only
echo 4. Run Both (requires 2 windows)
echo 5. Exit
echo.
set /p choice=Enter your choice (1-5): 

if "%choice%"=="1" goto full_setup
if "%choice%"=="2" goto run_agent
if "%choice%"=="3" goto run_dashboard
if "%choice%"=="4" goto run_both
if "%choice%"=="5" exit /b 0
echo Invalid choice
timeout /t 2 >nul
goto menu

:full_setup
echo.
echo Running full setup...
echo This will install all required dependencies.
echo.
call full-setup.bat
goto menu

:run_agent
echo.
echo Starting Agent...
echo.
cd agent
call setup-and-run.bat
cd ..
goto menu

:run_dashboard
echo.
echo Starting Dashboard...
echo.
cd web
call npm-setup-and-run.bat
cd ..
goto menu

:run_both
echo.
echo ================================================
echo Running Both Agent and Dashboard
echo ================================================
echo.
echo This will open 2 new windows...
echo.
timeout /t 2 >nul

REM Start agent in new window
start "APC Agent" cmd /k "cd agent && setup-and-run.bat"

REM Wait a moment
timeout /t 3 >nul

REM Start dashboard in new window
start "APC Dashboard" cmd /k "cd web && npm-setup-and-run.bat"

echo.
echo Both services started in new windows!
echo Open http://localhost:5173 in your browser
echo.
pause
goto menu

endlocal
