DOCKER_IMAGE := packing-service:latest
MAIN_PACKAGE := ./cmd/main
TEST_PACKAGE := ./...
DOCKERFILE := Dockerfile

.PHONY: all
all: build

.PHONY: build
build:
	go build -o main $(MAIN_PACKAGE)/main.go

.PHONY: test
test:
	go test $(TEST_PACKAGE) -v

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) -f $(DOCKERFILE) .

.PHONY: clean
clean:
	rm -f main