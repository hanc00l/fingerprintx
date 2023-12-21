#!/usr/bin/env bash

echo "build for macos,linux and windows..."

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o fingerprintx_darwin_amd64 cmd/fingerprintx/fingerprintx.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o fingerprintx_linux_amd64 cmd/fingerprintx/fingerprintx.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o fingerprintx_windows_amd64.exe cmd/fingerprintx/fingerprintx.go


echo "build done."