.PHONY: help
.DEFAULT_GOAL := help

help:
	@echo "---------------------------------------------------------------------------------------"
	@echo ""
	@echo "				CLI"
	@echo ""
	@echo "---------------------------------------------------------------------------------------"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

go_install: ## Install Go Dependencies
	go mod tidy && go mod download

go_build: go_install ## Build Go Dependencies
	go build -o build/anyops ./

go_release: go_build ## Release Go Binary
	rm -rf ~/bin/anyops && mkdir -p ~/bin/ && cp -R build/anyops ~/bin/
