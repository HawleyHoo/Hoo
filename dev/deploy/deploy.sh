#!/usr/bin/env bash

ID=`ps -ef | grep "testmain" | grep -v "$0" | grep -v "grep" | awk '{print $2}'`
echo $ID
echo "---------------"
echo $(pgrep testmain)
kill $(pgrep testmain)
cd /Users/lx/Desktop/goworks/src/Hoo/DevOps
#git pull git@gitlab.com:EdisonLeung/devops.git
go build -o "testmain" main.go
ls -l
nohup ./testmain & echo $! > pidfile.txt
echo $!
#kill -9 `cat pidfile.txt`
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o "${robot_name}" robot.go