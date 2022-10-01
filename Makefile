

.PHONY: build
build: clean
	CGO_ENABLED=0 go build -o output/ssh-key-sync

.PHONY: clean
clean:
	rm -f output/ssh-key-sync

.PHONY: test
test:
	go test -race ./...

.PHONY: vet
vet:
	go vet ./...

default: build
