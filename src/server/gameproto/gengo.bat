@ECHO OFF
protoc --go_out=. --proto_path=.;../../../../../../ gamedef/*.proto
protoc --go_out=. --proto_path=.;../../../../../../ cmsg/*.proto

ECHO.
ECHO Compile .proto To .go Done!
@IF %ERRORLEVEL% NEQ 0 PAUSE