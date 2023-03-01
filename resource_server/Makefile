IMAGE_TAG=suma-rs

.PHONY: build
build:
	@if docker images inspect ${IMAGE_TAG}; then \
		docker rmi ${IMAGE_TAG}; \
	fi
	docker build . -t ${IMAGE_TAG}
run:
	docker run --rm -p 127.0.0.1:10001:10001 --name suma-rs ${IMAGE_TAG}