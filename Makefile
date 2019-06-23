# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# Execute the `build` target by default.
.DEFAULT_GOAL := build

kind-deploy:
	@echo "> Deploying via Kubernetes in Docker..."
	go build -mod=vendor \
		-o tools/kind \
		sigs.k8s.io/kind
	PATH=$(PWD)/tools:$(PATH) scripts/kind-deploy.sh

proto-gen:
	@echo "> Generating Go protobuf files..."
	go build -mod=vendor \
		-o tools/protoc-gen-go \
		github.com/golang/protobuf/protoc-gen-go
	PATH=$(PWD)/tools:$(PATH) scripts/protoc-gen.sh

go-install:
	@echo "> Installing binaries into ${GOPATH} ..."
	go build -mod=vendor ./...

go-build:
	@echo "> Building all Go binaries here..."
	GOBIN=$(PWD) go install -mod=vendor ./...

go-compile:
	@echo "> Running the Go compiler..."
	go build -a -v -mod=vendor ./...

go-clean:
	@echo "> Cleaning the build cache..."
	go clean -mod=vendor ./...

go-test:
	@echo "> Running the tests ..."
	go test -v -count=1 -mod=vendor ./...

go-mod:
	@echo "> Upgrading all dependencies..."
	go mod tidy
	go mod vendor

## build: Build all Go binaries.
.PHONY: build
build: go-clean go-build

## install: Install Go binaries into GOPATH.
.PHONY: install
install: go-clean go-install

## compile: Run the Go compiler.
.PHONY: compile
compile: go-clean go-compile

## test: Run tests.
.PHONY: test
test: go-test

## upgrade: Upgrade and vendor dependencies.
.PHONY: upgrade
upgrade: go-mod

## proto: Generate protobuf files.
.PHONY: proto
proto: proto-gen

## docker-build: Build calc-server Docker image.
.PHONY: docker-build
docker-build:
	@echo "> Building in docker..."
	docker build \
		-t calc-server:v0 \
		-f build/package/Dockerfile .

## docker-run: Run calc-server Docker image.
.PHONY: docker-run
docker-run:
	docker run -it \
		--env-file=envfile \
		-p 9111:9111 \
		calc-server:v0

## deploy: Deploy calc-server on Kubernetes IN Docker.
.PHONY: deploy
deploy: kind-deploy

## clean: Clean build cache, remove binaries.
.PHONY: clean
clean: go-clean
	rm -f ./calc-server
	rm -f ./calc-client

.PHONY: help
help: Makefile
	@echo " Choose a make command:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
