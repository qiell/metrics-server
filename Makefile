BUILD_TAG:=$(shell git rev-parse --short HEAD)
ifeq ($(BRANCH),)
BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
endif

build: binary images

images: metrics-server-image
	@echo "Built all image"

metrics-server-image:
	docker build -f=build/Dockerfile.metrics-server -t qiell/metrics-server:$(BRANCH)-$(BUILD_TAG) .

binary: metrics-server
	@echo "Built all binaries"

metrics-server:
	@echo "Building metrics server"
	./scripts/build-server

.PHONY=build images metrics-server-image binary metrics-server