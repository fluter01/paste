# Copyright 2016 fluter

BIN=gopaste
DEST?=/usr/local/bin/

all: $(BIN)

$(BIN): paste.go
	go build -o $(BIN) cmd/gopaste/paste.go

install: $(BIN)
	install -D $(BIN) $(DEST)

test: paste_test.go
	go test -v -cover -covermode atomic -coverprofile coverage.out

.PHONY=clean
clean:
	rm -f $(BIN)
