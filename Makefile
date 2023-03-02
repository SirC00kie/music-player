APP_NAME := music-player
APP_PATH := ./cmd/$(APP_NAME)
TEST_PATH := ./tests
DOCKER_IMAGE := $(APP_NAME):latest

.PHONY: build
build:
	go build -o $(APP_NAME) $(APP_PATH)

.PHONY: run
run:
	go run $(APP_PATH)

.PHONY: test
test:
	go test $(TEST_PATH)/...

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run:
	docker run -it --rm $(DOCKER_IMAGE)

.PHONY: docker-test
docker-test:
	docker run --rm $(DOCKER_IMAGE) go test $(TEST_PATH)/...

.PHONY: clean
clean:
	rm -f $(APP_NAME)

.PHONY: clean-docker
clean-docker:
	docker rmi -f $(DOCKER_IMAGE)