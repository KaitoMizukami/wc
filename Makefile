BINARY_NAME=ccwc

all:build

build:
	go build -o $(BINARY_NAME) .

clean:
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME) $(ARGS)

test:
	go test