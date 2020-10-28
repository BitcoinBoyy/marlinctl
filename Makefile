GO=go
GOBUILD=$(GO) build
BINDIR=bin
BINCLI=cli
BIN=$(BINDIR)/$(BINCLI)

all:
	$(GOBUILD) -o $(BIN)
