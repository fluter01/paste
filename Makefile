
BIN=gopaste
DEST?=/usr/local/bin/

all: $(BIN)

$(BIN): paste.go
	go build -o $(BIN) cmd/paste.go

install: $(BIN)
	install -D $(BIN) $(DEST)

test: paste_test.go
	go test -v ./...

.PHONY=clean
clean:
	rm -f $(BIN)
