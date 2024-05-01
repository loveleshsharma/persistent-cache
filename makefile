test:
	@go test ./...

build-and-run:
	@echo "Running persistent-cache server..."
	@go build -C cmd/pcache && ./cmd/pcache/pcache