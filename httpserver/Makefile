.PHONY: build
build:
	@go build

.PHONY: run
run:
	@go run main.go

.PHONY: docker-build
docker-build:
	@docker build -t tenchael.com/httpserver .

.PHONY: docker-run
docker-run:
	@docker run -p 8080:8080 tenchael.com/httpserver
