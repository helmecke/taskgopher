# Build binary
bin: test
	go build -o build/taskgopher main.go

install: test
	go install

run:
	go run main.go

# Run tests against code
test: fmt vet
	go test ./...

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...
