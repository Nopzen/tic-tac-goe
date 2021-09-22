# Build tools
clean:
	rm -rf ./bin
.PHONY:

build: clean
	go build -o ./bin/tic-tac-goe ./cmd
.PHONY:

# Go Specfic targets
go_vet:
	go vet ./...
.PHONY:

go_get:
	go mod download
.PHONY:

go_test: go_vet
	go test ./...
.PHONY: