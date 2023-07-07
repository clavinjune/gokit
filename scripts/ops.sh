#!/bin/sh

lint_() {
  go run "$GOIMPORTS" -w .
  gofmt -w -s .
  go mod tidy
  go vet ./...
  go run "$GOVULNCHECK" ./...
}

test_() {
  go test -v -covermode=atomic -shuffle=on ./...
}

test_report_() {
  go test -covermode=atomic -shuffle=on -coverprofile coverage.out -json ./... > test-report.json
}

test_cover_() {
  go tool cover -html=coverage.out
}

set -euo pipefail

for mod in $(find . -name "go.mod"); do
  pushd "$(echo "$mod" | sed 's/\/go\.mod//g')"

  case $1 in
  lint)
    lint_
    ;;
  test)
    test_
    ;;
  test_report)
    test_report_
    ;;
  test_cover)
    test_cover_
    ;;
  esac

  popd
done

go work sync