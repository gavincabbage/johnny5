# johnny5 web

include ../Makefile.inc

SRC=main.go
BINARY=../$(TARGET_DIR)/web/web
GRUNT=grunt dist

web: $(SRC)
	$(GO) $(BINARY) $(SRC)
	$(GRUNT)

all: web
