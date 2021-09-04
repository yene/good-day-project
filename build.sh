#!/bin/bash
set -e
cd frontend
npm run build
cd ..
# GOARCH amd64,arm
GOOS=darwin go build -o good-day-macOS
upx good-day-macOS
GOOS=windows go build -o good-day.exe
upx good-day.exe
GOOS=linux go build -o good-day-linux
upx good-day-linux
