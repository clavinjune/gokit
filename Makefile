include tools.mk
export

.PHONY: lint
lint:
	@./scripts/ops.sh lint

.PHONY: test
test:
	@./scripts/ops.sh test

.PHONY: test/report
test/report:
	@./scripts/ops.sh test_report

.PHONY: test/cover
test/cover: test/report
	@./scripts/ops.sh test_cover
