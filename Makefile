include tools.mk

.PHONY: lint
lint:
	@go vet ./...
	@go run $(GOIMPORTS) -w .
	@gofmt -w -s .
	@go mod tidy
	@go run $(GOVULNCHECK) ./...


.PHONY: test
test:
	@go test -v -covermode=atomic -shuffle=on ./...

.PHONY: test/report
test/report:
	@go test -covermode=atomic -shuffle=on -coverprofile coverage.out -json ./... > test-report.json

.PHONY: test/cover
test/cover: test/report
	@go tool cover -html=coverage.out
