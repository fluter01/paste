
BIN=gopaste
DEST?=/usr/local/bin/

all: $(BIN)

$(BIN):
	go build -o $(BIN) cmd/paste.go

install: $(BIN)
	install -D $(BIN) $(DEST)

.PHONY=clean
clean:
	rm -f $(BIN)
