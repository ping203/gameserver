#!/bin/bash
protoc --go_out=. --proto_path=.\;../../../../ gamedef/*.proto
protoc --go_out=. --proto_path=.\;../../../../ cmsg/*.proto
protoc --go_out=. --proto_path=.\;../../../../ smsg/*.proto
protoc --go_out=. --proto_path=.\;../../../../ wmsg/*.proto
protoc --go_out=. --proto_path=.\;../../../../ emsg/*.proto
protoc --go_out=. --proto_path=.\;../../../../ netframe/*.proto
protoc --go_out=. --proto_path=.\;../../../../ logicmsg/*.proto
protoc --go_out=. --proto_path=.\;../../../../ gameconf/*.proto

echo "."
echo "Compile .proto To .go Done!"