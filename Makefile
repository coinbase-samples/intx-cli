.PHONY: build install

BINARY_NAME=intxctl

all: build

build:
	go build -o $(BINARY_NAME)

install:
	sudo mv $(BINARY_NAME) /usr/local/bin/
	sudo chmod 755 /usr/local/bin/$(BINARY_NAME)

path-note:
	@echo "If the binary is not found, you may need to add /usr/local/bin to your \$PATH."