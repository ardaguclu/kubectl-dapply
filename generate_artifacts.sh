#!/bin/bash

env GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build cmd/kubectl-dapply.go && tar -zcvf kubectl-dapply_v0.24.3_darwin_arm64.tar.gz kubectl-dapply LICENSE
sha256sum kubectl-dapply_v0.24.3_darwin_arm64.tar.gz

env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build cmd/kubectl-dapply.go && tar -zcvf kubectl-dapply_v0.24.3_darwin_amd64.tar.gz kubectl-dapply LICENSE
sha256sum kubectl-dapply_v0.24.3_darwin_amd64.tar.gz

env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build cmd/kubectl-dapply.go && tar -zcvf kubectl-dapply_v0.24.3_linux_amd64.tar.gz kubectl-dapply LICENSE
sha256sum kubectl-dapply_v0.24.3_linux_amd64.tar.gz

env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build cmd/kubectl-dapply.go && tar -zcvf kubectl-dapply_v0.24.3_windows_amd64.tar.gz kubectl-dapply.exe LICENSE
sha256sum kubectl-dapply_v0.24.3_windows_amd64.tar.gz
