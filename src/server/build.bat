@ECHO OFF

@REM compile and build
go build -o ../../bin/server.exe ./

@REM copy game config
xcopy /Y .\gameproto\gameconf\*.pbt ..\..\bin\gameconf\

@IF %ERRORLEVEL% NEQ 0 PAUSE
