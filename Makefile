.PHONY: build run test init-test-config

#build:
#	go build -o ./bin/app ./cmd/app/main.go
#
#run: build
#	./bin/app --config=./config/config.yaml
#
#test:
#	go test -v ./tests

SECRETS_DIR=./tests/test-config/secrets
LOGS_DIR=./tests/test-config/logs
GITHUB_TOKEN_FILE=$(SECRETS_DIR)/github_token
AUTH_SERVICE_LOG_FILE=$(LOGS_DIR)/auth-service.log
DOCKER_COMPOSE_FILE=./tests/test-config/docker-compose.yaml

init-test-config:
	@echo "Checking for required files..."

	@mkdir -p $(SECRETS_DIR)
	@mkdir -p $(LOGS_DIR)

	@if ! grep -q "ghp" $(GITHUB_TOKEN_FILE); then \
    		echo "github_token found, but it does not contain 'ghp'."; \
    		echo "Please enter your GitHub Token (starting with 'ghp'):"; \
    		read -r TOKEN; \
    		if [ -n "$$TOKEN" ]; then \
    			echo "$$TOKEN" > $(GITHUB_TOKEN_FILE); \
    			echo "GitHub Token saved to $(GITHUB_TOKEN_FILE)."; \
    		else \
    			echo "No input provided. GitHub Token not updated."; \
    		fi; \
    	else \
    		echo "github_token already contains a valid key."; \
    	fi

	@if [ ! -f $(AUTH_SERVICE_LOG_FILE) ]; then \
		echo "auth-service.log not found. Creating it..."; \
		touch $(AUTH_SERVICE_LOG_FILE); \
	else \
		echo "auth-service.log already exists."; \
	fi

build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

run: build
	docker-compose -f $(DOCKER_COMPOSE_FILE) up