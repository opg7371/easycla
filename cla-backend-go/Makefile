# Copyright The Linux Foundation and each contributor to CommunityBridge.
# SPDX-License-Identifier: MIT
SERVICE = cla
SHELL = bash
BIN_DIR = bin
LAMBDA_BIN = backend-aws-lambda
METRICS_BIN = metrics-aws-lambda
METRICS_REPORT_BIN = metrics-report-lambda
DYNAMO_EVENTS_BIN = dynamo-events-lambda
ZIPBUILDER_SCHEDULER_BIN = zipbuilder-scheduler-lambda
ZIPBUILDER_BIN = zipbuilder-lambda
GITLAB_REPO_CHECK_BIN = gitlab-repository-check-lambda
FUNCTIONAL_TESTS_BIN = functional-tests
USER_SUBSCRIBE_BIN = user-subscribe-lambda
MAKEFILE_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BUILD_TIME=$(shell sh -c 'date -u +%FT%T%z')
VERSION := $(shell sh -c 'git describe --always --tags')
BRANCH := $(shell sh -c 'git rev-parse --abbrev-ref HEAD')
COMMIT := $(shell sh -c 'git rev-parse --short HEAD')
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.branch=$(BRANCH) -X main.buildDate=$(BUILD_TIME)"
BUILD_TAGS=-tags aws_lambda

ifeq "$(shell uname -s)" "Darwin"
	BUILD_HOST=darwin
endif
ifeq "$(shell uname -s)" "Linux"
	BUILD_HOST=linux
endif

LINT_TOOL=$(shell go env GOPATH)/bin/golangci-lint
LINT_VERSION=v1.46.2
SWAGGER_DIR=$(ROOT_DIR)/swagger
SWAGGER_BIN_DIR=/usr/local/bin
SWAGGER_TOOL_VERSION=v0.24.0
SWAGGER_ASSET="swagger_$(BUILD_HOST)_amd64"
SWAGGER_ASSET_URL="https://github.com/go-swagger/go-swagger/releases/download/$(SWAGGER_TOOL_VERSION)/$(SWAGGER_ASSET)"
GO_PKGS=$(shell go list ./... | grep -v /vendor/ | grep -v /node_modules/)
GO_FILES=$(shell find . -type f -name '*.go' -not -path './vendor/*')

.PHONY: generate setup setup-dev setup-deploy clean-all clean swagger up fmt test run deps build build-mac build-aws-lambda user-subscribe-lambda qc lint

all: all-mac
all-mac: clean swagger deps fmt build-mac build-aws-lambda-mac build-user-subscribe-lambda-mac build-metrics-lambda-mac build-dynamo-events-lambda-mac build-zipbuilder-scheduler-lambda-mac build-zipbuilder-lambda-mac build-gitlab-repository-check-lambda-mac test lint
all-linux: clean swagger deps fmt build-linux build-aws-lambda-linux build-user-subscribe-lambda-linux build-metrics-lambda-linux build-dynamo-events-lambda-linux build-zipbuilder-scheduler-lambda-linux build-zipbuilder-lambda-linux build-gitlab-repository-check-lambda-linux test lint
lambdas-mac: build-lambdas-mac
build-lambdas-mac: build-aws-lambda-mac build-user-subscribe-lambda-mac build-metrics-lambda-mac build-metrics-report-lambda-mac build-dynamo-events-lambda-mac build-zipbuilder-scheduler-lambda-mac build-zipbuilder-lambda-mac build-gitlab-repository-check-lambda-mac
lambdas: build-lambdas-linux
build-lambdas-linux: build-aws-lambda-linux build-user-subscribe-lambda-linux build-metrics-lambda-linux build-metrics-report-lambda-linux build-dynamo-events-lambda-linux build-zipbuilder-scheduler-lambda-linux build-zipbuilder-lambda-linux build-gitlab-repository-check-lambda-linux

generate: swagger

setup: $(LINT_TOOL) setup-dev setup-deploy

.PHONY: swagger-tool
swagger-tool:
	@echo "Removing old swagger binary in $(SWAGGER_BIN_DIR)..."
	@sudo rm -Rf $(SWAGGER_BIN_DIR)/swagger
	@echo "Downloading $(SWAGGER_ASSET_URL)"
	sudo curl -o $(SWAGGER_BIN_DIR)/swagger -L'#' $(SWAGGER_ASSET_URL)
	sudo chmod +x $(SWAGGER_BIN_DIR)/swagger
	$(SWAGGER_BIN_DIR)/swagger version

setup_dev: setup-dev
setup-dev: swagger-tool
	pushd /tmp && echo "Installing goimport..." && go install golang.org/x/tools/cmd/goimports@latest && echo "Installation coverage tools..." && go install golang.org/x/tools/cmd/cover@latest && popd

	@echo "Installing linter..."
	@# Latest releases: https://github.com/golangci/golangci-lint/releases
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINT_VERSION)
	echo "golangci-lint version:" && golangci-lint version

	@echo "Installing multi-file-swagger tool..."
	cd $(dir $(realpath $(firstword $(MAKEFILE_LIST))))swagger && pip3 install virtualenv && virtualenv .venv && source .venv/bin/activate && python -m pip install --upgrade pip && pip3 install -r requirements.txt

setup_deploy: setup-deploy
setup-deploy:
	@yarn add serverless && yarn install

clean: clean-models clean-lambdas
	@rm -rf cla cla-mac* cla-linux

clean-models:
	@rm -rf ./v2/project-service/client ./v2/project-service/models \
		./v2/organization-service/client ./v2/organization-service/models \
		./v2/user-service/client ./v2/user-service/models

clean-lambdas:
	@rm -rf $(BIN_DIR)

swagger-clean: clean-swagger
clean-swagger:
	@rm -rf gen/

clean-all: clean clean-swagger
	@rm -rf vendor/

swagger: clean-swagger swagger-prep swagger-build swagger-validate
build-swagger: swagger-build
swagger-build: swagger-build-v1-services swagger-build-v2-services swagger-build-project-service swagger-build-organization-service swagger-build-user-service swagger-build-acs-service
swagger-validate: swagger-v1-validate swagger-v2-validate

swagger-prep:
	@mkdir gen

swagger-build-v1-services:
	@echo
	@echo "Generating v1 legacy API models..."
	cd swagger; source .venv/bin/activate && python3 multi-file-swagger.py --spec-input-file cla.v1.yaml --spec-output-file cla.v1.compiled.yaml
	swagger -q generate server \
		-t gen \
		-f swagger/cla.v1.compiled.yaml \
		--copyright-file=copyright-header.txt \
		--server-package=v1/restapi \
		--model-package=v1/models \
		--exclude-main \
		-A cla \
		-P user.CLAUser

swagger-build-v2-services:
	@echo
	@echo "Generating v2 API models..."
	cd swagger; source .venv/bin/activate && python3 multi-file-swagger.py --spec-input-file cla.v2.yaml --spec-output-file cla.v2.compiled.yaml
	swagger -q generate server \
		-t gen \
		-f swagger/cla.v2.compiled.yaml \
		--copyright-file=copyright-header.txt \
		--server-package=v2/restapi \
		--model-package=v2/models \
		--exclude-main \
		-A easycla \
		-P auth.User

swagger-build-project-service:
	@echo
	@echo "Generating swagger client for the project-service..."
	@mkdir -p v2/project-service
	curl -sfL https://api-gw.dev.platform.linuxfoundation.org/project-service/swagger.json --output swagger/project-service.yaml
	sed -i.bak 's/X-ACL/Empty-Header/g' swagger/project-service.yaml
	swagger -q generate client \
		--copyright-file=copyright-header.txt \
		-t v2/project-service \
		-f swagger/project-service.yaml \
		--skip-validation # needed, currently seeing: body.default.Filename in body must be of type string: "null", and definitions.artifact-upload-init-request.default.Filename in body must be of type string: "null" issues, notified PS team

swagger-build-organization-service:
	@echo
	@echo "Generating swagger client for the organization-service..."
	@mkdir -p v2/organization-service
	curl -sfL https://api-gw.dev.platform.linuxfoundation.org/organization-service/swagger.json --output swagger/organization-service.yaml
	sed -i.bak 's/X-ACL/Empty-Header/g' swagger/organization-service.yaml
	swagger -q generate client \
		--copyright-file=copyright-header.txt \
		-t v2/organization-service \
		-f swagger/organization-service.yaml \
		--skip-validation # needed, currently seeing: - username in query must be of type string: "null"

swagger-build-user-service:
	@echo
	@echo "Generating swagger client for the user-service..."
	@mkdir -p v2/user-service
	curl -sfL https://api-gw.dev.platform.linuxfoundation.org/user-service/swagger.json --output swagger/user-service.yaml
	sed -i.bak 's/X-ACL/Empty-Header/g' swagger/user-service.yaml
	swagger -q generate client \
		--copyright-file=copyright-header.txt \
		-t v2/user-service \
		-f swagger/user-service.yaml \
		--skip-validation # needed, many validation errors

swagger-build-acs-service:
	@echo
	@echo "Generating swagger client for the acs-service..."
	@mkdir -p v2/acs-service
	curl -sfL https://api-gw.dev.platform.linuxfoundation.org/acs/v1/api-docs/swagger/swagger.json --output swagger/acs-service.yaml
	sed -i.bak 's/X-ACL/X-API-KEY/g' swagger/acs-service.yaml
	swagger -q generate client \
		--copyright-file=copyright-header.txt \
		-t v2/acs-service \
		-f swagger/acs-service.yaml

swagger-v1-validate:
	@echo ""
	@echo "Validating EasyCLA v1 legacy API specification..."
	@swagger validate --stop-on-error swagger/cla.v1.compiled.yaml

swagger-v2-validate:
	@echo ""
	@echo "Validating EasyCLA v2 API specification..."
	@swagger validate --stop-on-error swagger/cla.v2.compiled.yaml

fmt:
	@echo "Formatting code and optimizing imports..."
	@gofmt -w -l -s $(GO_FILES)
	@goimports -w -l $(GO_FILES)

test:
	@echo "Running unit tests..."
	@go test -v $(shell go list ./... | grep -v /vendor/ | grep -v /node_modules/) -coverprofile=cover.out

mock:
	@echo "Re-Generating mocks"
	@cd $(MAKEFILE_DIR) && mkdir -p repositories/mock
	@cd $(MAKEFILE_DIR) && mockgen -copyright_file=copyright-header.txt -source=repositories/service.go -package=mock -destination=repositories/mock/mock_service.go
	@cd $(MAKEFILE_DIR) && mockgen -copyright_file=copyright-header.txt -source=repositories/repository.go -package=mock -destination=repositories/mock/mock_repository.go
    # github organizations
	@cd $(MAKEFILE_DIR) && mockgen -copyright_file=copyright-header.txt -package=github_organizations -destination=github_organizations/mock.go -self_package=github.com/communitybridge/easycla/cla-backend-go/github_organizations github.com/communitybridge/easycla/cla-backend-go/github_organizations Repository
    # events
	@cd $(MAKEFILE_DIR) && mockgen -copyright_file=copyright-header.txt -package=events -destination=events/mock.go -self_package=github.com/communitybridge/easycla/cla-backend-go/events github.com/communitybridge/easycla/cla-backend-go/events Service
	# mocks for github
	@cd $(MAKEFILE_DIR) && mockgen -copyright_file=copyright-header.txt -package=branch_protection -destination=github/branch_protection/mock.go -self_package=github.com/communitybridge/easycla/cla-backend-go/github/branch_protection github.com/communitybridge/easycla/cla-backend-go/github/branch_protection CombinedRepository


deps:
	go env -w GOPRIVATE=github.com/LF-Engineering/*
	go mod download -x

build-prep:
	mkdir -p $(BIN_DIR)

build: build-linux
build-linux: deps build-prep
	@echo "Building Linux amd64 binary..."
	env GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(SERVICE) main.go
	@chmod +x $(BIN_DIR)/$(SERVICE)

build-mac: deps build-prep
	@echo "Building Mac OSX amd64 binary..."
	env GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(SERVICE)-mac main.go
	@chmod +x $(BIN_DIR)/$(SERVICE)-mac

build-mac-arm64: deps build-prep
	@echo "Building Mac OSX Arm64 binary..."
	env GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BIN_DIR)/$(SERVICE)-mac-arm64 main.go
	@chmod +x $(BIN_DIR)/$(SERVICE)-mac-arm64

rebuild-mac: clean fmt build-mac lint

build-aws-lambda: build-aws-lambda-linux
build-aws-lambda-linux: deps build-prep
	@echo "Building a statically linked Linux amd64 binary..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) $(BUILD_TAGS) -o $(BIN_DIR)/$(LAMBDA_BIN) main.go
	@chmod +x $(BIN_DIR)/$(LAMBDA_BIN)

build-aws-lambda-mac: deps build-prep
	@echo "Building a statically linked Mac OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) $(BUILD_TAGS) -o $(BIN_DIR)/$(LAMBDA_BIN)-mac main.go
	@chmod +x $(BIN_DIR)/$(LAMBDA_BIN)-mac

build-user-subscribe-lambda: build-user-subscribe-lambda-linux
build-user-subscribe-lambda-linux: deps build-prep
	@echo "Building a statically linked Linux amd64 binary..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) $(BUILD_TAGS) -o $(BIN_DIR)/$(USER_SUBSCRIBE_BIN) cmd/user-subscribe-lambda/main.go
	@chmod +x $(BIN_DIR)/$(USER_SUBSCRIBE_BIN)

build-user-subscribe-lambda-mac: deps build-prep
	@echo "Building a statically linked Mac OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) $(BUILD_TAGS) -o $(BIN_DIR)/$(USER_SUBSCRIBE_BIN)-mac cmd/user-subscribe-lambda/main.go
	@chmod +x $(BIN_DIR)/$(USER_SUBSCRIBE_BIN)-mac

build-metrics-lambda: build-metrics-lambda-linux
build-metrics-lambda-linux: deps build-prep
	@echo "Building a statically linked Linux amd64 binary..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(METRICS_BIN) cmd/metrics_lambda/main.go
	@chmod +x $(BIN_DIR)/$(METRICS_BIN)

build-metrics-lambda-mac: deps build-prep
	@echo "Building a statically linked Mac OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(METRICS_BIN)-mac cmd/metrics_lambda/main.go
	@chmod +x $(BIN_DIR)/$(METRICS_BIN)-mac

build-metrics-report-lambda: build-metrics-report-lambda-linux
build-metrics-report-lambda-linux: deps build-prep
	@echo "Building a statically linked Linux amd64 binary..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(METRICS_REPORT_BIN) cmd/metrics_report_lambda/main.go
	@chmod +x $(BIN_DIR)/$(METRICS_REPORT_BIN)

build-metrics-report-lambda-mac: deps build-prep
	@echo "Building a statically linked Mac OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(METRICS_REPORT_BIN)-mac cmd/metrics_report_lambda/main.go
	@chmod +x $(BIN_DIR)/$(METRICS_REPORT_BIN)-mac

build-dynamo-events-lambda: build-dynamo-events-lambda-linux
build-dynamo-events-lambda-linux: deps build-prep
	@echo "Building a statically linked Linux amd64 binary..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(DYNAMO_EVENTS_BIN) cmd/dynamo_events_lambda/main.go
	@chmod +x $(BIN_DIR)/$(DYNAMO_EVENTS_BIN)

build-dynamo-events-lambda-mac: deps build-prep
	@echo "Building a statically linked Mac OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(DYNAMO_EVENTS_BIN)-mac cmd/dynamo_events_lambda/main.go
	@chmod +x $(BIN_DIR)/$(DYNAMO_EVENTS_BIN)-mac

build-zipbuilder-scheduler-lambda: build-zipbuilder-scheduler-lambda-linux
build-zipbuilder-scheduler-lambda-linux: deps build-prep
	@echo "Building a statically linked Linux amd64 binary..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(ZIPBUILDER_SCHEDULER_BIN) cmd/zipbuilder_scheduler_lambda/main.go
	@chmod +x $(BIN_DIR)/$(ZIPBUILDER_SCHEDULER_BIN)

build-zipbuilder-scheduler-lambda-mac: deps build-prep
	@echo "Building a statically linked Mac OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(ZIPBUILDER_SCHEDULER_BIN)-mac cmd/zipbuilder_scheduler_lambda/main.go
	@chmod +x $(BIN_DIR)/$(ZIPBUILDER_SCHEDULER_BIN)-mac

build-zipbuilder-lambda: build-zipbuilder-lambda-linux
build-zipbuilder-lambda-linux: deps build-prep
	@echo "Building a statically linked Linux amd64 binary..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(ZIPBUILDER_BIN) cmd/zipbuilder_lambda/main.go
	@chmod +x $(BIN_DIR)/$(ZIPBUILDER_BIN)

build-zipbuilder-lambda-mac: deps build-prep
	@echo "Building a statically linked Mac OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(ZIPBUILDER_BIN)-mac cmd/zipbuilder_lambda/main.go
	@chmod +x $(BIN_DIR)/$(ZIPBUILDER_BIN)-mac

build-gitlab-repository-check-lambda-linux: deps build-prep
	@echo "Building a statically linked Linux OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) $(BUILD_TAGS) -o $(BIN_DIR)/$(GITLAB_REPO_CHECK_BIN) cmd/gitlab_repository_check/main.go
	@chmod +x $(BIN_DIR)/$(GITLAB_REPO_CHECK_BIN)

build-gitlab-repository-check-lambda-mac: deps build-prep
	@echo "Building a statically linked Mac OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(GITLAB_REPO_CHECK_BIN)-mac cmd/gitlab_repository_check/main.go
	@chmod +x $(BIN_DIR)/$(GITLAB_REPO_CHECK_BIN)-mac

build-functional-tests: build-functional-tests-linux
build-functional-tests-linux: deps build-prep
	@echo "Building Functional Tests for Linux amd64 binary..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(FUNCTIONAL_TESTS_BIN) cmd/functional_tests/main.go
	@chmod +x $(BIN_DIR)/$(FUNCTIONAL_TESTS_BIN)

build-functional-tests-mac: deps build-prep
	@echo "Building Functional Tests for OSX amd64 binary..."
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(FUNCTIONAL_TESTS_BIN)-mac cmd/functional_tests/main.go
	@chmod +x $(BIN_DIR)/$(FUNCTIONAL_TESTS_BIN)-mac

lint:
	@cd $(MAKEFILE_DIR) && echo "Running lint..." && $(LINT_TOOL) --version && $(LINT_TOOL) run --exclude="this method will not auto-escape HTML. Verify data is well formed" --allow-parallel-runners --config=.golangci.yaml ./... && echo "Lint check passed."
	@cd $(MAKEFILE_DIR) && ./check-headers.sh

