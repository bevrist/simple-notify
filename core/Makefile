.PHONY: help
help: # prints self documentation
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Compile and run the program
	@go run ./

test: ## Run the tests
	@find . -name "*_test.go" -execdir go test ./ \;
  # @find . -name "*_test.go" -exec dirname {} \; | xargs -P8 go test

testv: ## Run the tests verbosely
	@find . -name "*_test.go" -execdir go test -v ./ \;

bench: ## run all benchmarks
	@find . -name "*_test.go" -execdir go test -benchmem -bench=. \;
