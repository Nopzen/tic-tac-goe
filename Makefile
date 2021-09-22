go_vet:
	go vet ./...
.PHONY:

go_test:
	go test ./...
.PHONY:

go_get:
	go mod download
.PHONY:

clean:
	rm -rf ./bin
.PHONY:

build: clean go_get go_vet go_test
	go build -o ./bin/tic-tac-goe ./cmd
.PHONY:
