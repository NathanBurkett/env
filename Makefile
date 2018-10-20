.FORCE:
.PHONY: help

help:
	@echo $(info Run any of the below commands as make {command})
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
	@echo ""

install: ## Install dependencies
	@dep ensure

lint:
	@$(GOPATH)/bin/golint -set_exit_status

html-coverage: ## Run tests with coverage and view html output
	@./coverage/run.sh --html

coverage: ## Run tests with coverage
	@./coverage/run.sh
