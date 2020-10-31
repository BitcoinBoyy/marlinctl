GO=go
GOBUILD=$(GO) build
BINDIR=bin
BINCLI=marlinctl
BIN=$(BINDIR)/$(BINCLI)

all:
	$(GOBUILD) -o $(BIN)
