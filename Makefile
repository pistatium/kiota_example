help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init:
	go install github.com/swaggo/swag/cmd/swag@latest

run_webapi: init  ## Run example webapi
	cd webapi && swag init
	go run ./webapi

generate_client: ./webapi/docs/swagger.yaml ## Generate client code via kiota
	kiota generate -l go -d ./webapi/docs/swagger.yaml -o ./client -n github.com/pistatium/kiota_example/client

run_example_client: ## Run example client
	go run ./...