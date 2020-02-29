# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=subtitlefixer
BINARY_UNIX=$(BINARY_NAME)_unix
all: clean build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	clear
	./$(BINARY_NAME)
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)
	rm -rf $(BINARY_UNIX)
	clear
dev:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME) -f "subtitle.srt"
	./$(BINARY_NAME) -f "subtitle.srt" -o "new subtitle.srt"
	./$(BINARY_NAME) --file "subtitle.srt" -output "new subtitle1.srt"

