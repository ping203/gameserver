@ECHO OFF
protoc --go_out=. --proto_path=.;../../../../../../ gamedef/*.proto
protoc --go_out=. --proto_path=.;../../../../../../ cmsg/*.proto
protoc --go_out=. --proto_path=.;../../../../../../ smsg/*.proto
protoc --go_out=. --proto_path=.;../../../../../../ emsg/*.proto
protoc --go_out=. --proto_path=.;../../../../../../ gameconf/*.proto

ECHO.
ECHO Compile .proto To .go Done!
@IF %ERRORLEVEL% NEQ 0 PAUSE