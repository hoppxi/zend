@echo off
SETLOCAL

:: Determine install directory
set PROGRAM_DIR=%ProgramFiles%\Zend
set APPDATA_DIR=%ProgramData%\Zend
set BIN_DIR=%PROGRAM_DIR%

:: Check if running as admin
>nul 2>&1 "%SYSTEMROOT%\system32\cacls.exe" "%SYSTEMROOT%\system32\config\system"
if '%errorlevel%' NEQ '0' (
    echo No admin privileges, switching to user install
    set PROGRAM_DIR=%LocalAppData%\Zend
    set APPDATA_DIR=%APPDATA%\Zend
    set BIN_DIR=%PROGRAM_DIR%
)

if not exist "%PROGRAM_DIR%" mkdir "%PROGRAM_DIR%"
if not exist "%APPDATA_DIR%" mkdir "%APPDATA_DIR%"

echo Building Go binary...
go build -o "%PROGRAM_DIR%\zend.exe" .\cmd

echo Building frontend...
cd web
npm install
npm run build
cd ..
xcopy /E /I /Y web\dist "%PROGRAM_DIR%\dist"

echo Writing default config...
if not exist "%APPDATA_DIR%\config.yaml" copy "%PROGRAM_DIR%\config.yaml" "%APPDATA_DIR%\config.yaml"

echo Setting ZEND_DIST environment variable...
setx ZEND_DIST "%PROGRAM_DIR%\dist"

echo Installation complete.
echo Binary: %BIN_DIR%\zend.exe
echo Config: %APPDATA_DIR%\config.yaml
echo Dist: %PROGRAM_DIR%\dist
ENDLOCAL
pause
