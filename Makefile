run:
	go run cmd/cli/main.go

build:
	go build -o bin/task-tracker cmd/cli/main.go

test:
	go test ./... -v
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

tidy:
	go mod download
	go mod tidy

clean:
	rm -fr bin
