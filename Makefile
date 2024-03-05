help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

init:
	go install github.com/swaggo/swag/cmd/swag@latest

	# OpenAPI 3.1 を生成する場合(gin-swaggerが未対応)
	# go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc3

run_webapi: init  ## Run example webapi
	cd webapi && swag init
	# OpenAPI 3.1 を生成する場合
	# cd webapi && swag init --v3.1
	go run ./webapi

generate_client: ./webapi/docs/swagger.yaml ## Generate client code via kiota
	kiota generate -l go -d ./webapi/docs/swagger.yaml -o ./client -n github.com/pistatium/kiota_example/client

run_example_client: ## Run example client
	go run ./...