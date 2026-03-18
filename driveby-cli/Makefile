.PHONY: build test vet clean

BINARY := driveby

build:
	go build -o $(BINARY) ./cmd/driveby

test:
	go test ./...

vet:
	go vet ./...

clean:
	rm -f $(BINARY)
