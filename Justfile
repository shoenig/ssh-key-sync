set shell := ["bash", "-u", "-c"]

export scripts := ".github/workflows/scripts"
export GOBIN := `echo $PWD/.bin`
export TAG := `git describe --tags $(git rev-list --tags --max-count=1)`

# print available commands
[private]
default:
    @just --list

# tidy up Go modules
[group('build')]
tidy:
    go mod tidy

# compile the executable
[group('build')]
compile: tidy
    go install

# run specific unit test
[group('build')]
[no-cd]
test unit:
    go test -v -count=1 -race -run {{unit}} 2>/dev/null

# run tests across source tree
[group('build')]
tests:
    go test -v -race -count=1 ./...

# vet the source tree
[group('lint')]
vet:
    go vet ./...

# lint the source tree
[group('lint')]
lint: vet
    $GOBIN/golangci-lint run --config {{scripts}}/golangci.yaml

# locally install build tools
[group('build')]
init:
    go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4

# show host system information
[group('build')]
@sysinfo:
    echo "{{os()/arch()}} {{num_cpus()}}c"

# publish a release
[group('release')]
release:
    export GORELEASER_CURRENT_TAG={{TAG}}
    envy exec gh-release goreleaser release --clean
    rm -rf dist
