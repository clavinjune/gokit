#!/bin/sh

lint() {
  pushd "$1"

  go mod tidy
  go vet ./...
  go run "$GOVULNCHECK" ./...
  popd
}

set -euo pipefail

go run "$GOIMPORTS" -w .
gofmt -w -s .

for mod in $(find . -name "go.mod"); do
  lint "$(echo "$mod" | sed 's/\/go\.mod//g')"
done

go work sync