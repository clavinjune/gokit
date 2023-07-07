include tools.mk

.PHONY: lint
lint: ./*util
	@go run $(GOIMPORTS) -w .
	@gofmt -w -s .
	for mod in $^ ; do \
		pushd ./$${mod} && \
		go mod tidy && \
		popd ; \
		go vet ./$${mod}/... ; \
		go run $(GOVULNCHECK) ./$${mod}/ ; \
	done

.PHONY: test
test: ./*util
	@for mod in $^ ; do \
		pushd ./$${mod} && \
		go test -v -covermode=atomic -shuffle=on ./... && \
		popd ; \
	done

.PHONY: test/report
test/report: ./*util
	@for mod in $^ ; do \
		pushd ./$${mod} && \
        go test -covermode=atomic -shuffle=on -coverprofile coverage.out -json ./... > test-report.json && \
        popd ; \
	done

#.PHONY: test/cover
#test/cover: test/report
#	@go tool cover -html=coverage.out
