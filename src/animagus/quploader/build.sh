#!/bin/bash
if [[ "$1" =~ "linux" ]] ;then     
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o q_upload_dir_linux .
else 
    go build -o q_upload_dir .
fi 

