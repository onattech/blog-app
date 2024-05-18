# Commands
.PHONY: list
list:
	@echo "📋 Available commands:"
	@awk -F':.*?## ' '/^[a-zA-Z0-9_-]+:/ && !/^[[:blank:]]*list:/ { if ($$2 == "") { printf "   • %s\n", $$1 } else { printf "   • %-20s %s\n", $$1, $$2 } }' $(MAKEFILE_LIST)

.PHONY: dev
dev: ## 📐 Starts the app in dev environment
	go run main.go

.PHONY: air 
air: ## 📐 Starts the app in dev environment with hot reload 🔥
	air

.PHONY: test
test: ## 🧪 Runs the tests
	go test -v ./...

.PHONY: testnocache
testnocache: ## 🧪 Runs the tests caching turned off
	go test -v -count=1  ./...

.PHONY: coverage
coverage: ## 📊 Displays test coverage
	go test -v ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

.PHONY: build
build: ## 🏗️  Builds the app
	go build
	
.PHONY: lint
lint: ## ✅ Checks for linting errors
	golangci-lint run -D errcheck
