NAME=tic-tac-goe

OUTPUT_DIR=bin

prefix?=/usr/local
mandir?=$(prefix)/share/man
bindir?=$(prefix)/bin

# Build tools
clean:
	rm -rf ./bin
.PHONY:

build: clean
	go build -o ./bin/$(NAME) ./cmd
.PHONY:

# Installation script
install:
	@@install -dm 755 $(bindir)/
	@@install -m 755 $(OUTPUT_DIR)/$(NAME) $(bindir)/$(NAME)
	@@mkdir -p $(mandir)/man1
	@@cp ./man/tic-tac-goe.1 $(mandir)/man1/tic-tac-goe.1
	@@mandb
	@@echo
	@@echo "-------------------------------------"
	@@echo "   Installation complete, enjoy!"
	@@echo "-------------------------------------"
.PHONY:

uninstall:
	rm -rf $(bindir)/$(NAME)
	rm -rf $(mandir)/man1/tic-tac-goe.1
	mandb
	@@echo
	@@echo "-------------------------------------------"
	@@echo "   Successfully uninstalled Tic Tac Goe    "
	@@echo "-------------------------------------------"
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