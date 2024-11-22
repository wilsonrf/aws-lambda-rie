#!/usr/bin/env bash
set -euo pipefail

GOMOD=$(head -1 go.mod | awk '{print $2}')
GOOS="linux" GOARCH="amd64" go build -ldflags='-s -w' -o linux/amd64/bin/main "$GOMOD/cmd/main"
GOOS="linux" GOARCH="amd64" go build -ldflags='-s -w' -o bin/main "$GOMOD/cmd/main"

ln -fs main linux/amd64/bin/generate
ln -fs main linux/amd64/bin/detect

ln -fs main bin/generate
ln -fs main bin/detect