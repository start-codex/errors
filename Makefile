.PHONY: test
test:
	@echo "Running tests..."
	cd $(folder) && go test -coverprofile=coverage.out && go tool cover -html=coverage.out
	@echo "Done."