build:
	@echo "Building droplet at './bin/droplet' ..."
	@go build -o bin/droplet cmd/droplet/*.go

all: lint	vet	cyclo	test	build

test:
	@echo "Running unit tests..."
	@go test -cover ./internal/... ./pkg/...

cyclo:
	@echo "Checking cyclomatic complexity..."
	@gocyclo -over 7 ./internal ./pkg

vet:
	@echo "Running vet..."
	@go vet ./...

lint:
	@echo "Running golint..."
	@golint ./internal/... ./pkg/...

setup:
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/fzipp/gocyclo