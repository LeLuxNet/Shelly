@echo off

call :build windows
call :build linux

echo.
cd ../out
echo Finished!
exit /B 0

:build
set GOOS=%~1
echo Build %GOOS%
go build -o ../out -ldflags "-s -w" ../...
exit /B 0