

.PHONY: build
build: clean
	CGO_ENABLED=0 go build -o output/ssh-key-sync

clean:
	rm -f output/ssh-key-sync

test:
	go test -race ./...

vet:
	go vet ./...

default: build
