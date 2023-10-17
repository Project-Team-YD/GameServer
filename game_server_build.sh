#!/bin/bash



echo "---------- Start GameServer Build ----------"
GOOS=linux GOARCH=amd64 go build -o game_server
echo "Buile Success"
echo "--------------------------------------------"