run:
	go run main.go

build:
	go build -o ./bin/task-tracker

test:
	go test ./... -v
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

tidy:
	go mod download
	go mod tidy

clean:
	rm -fr bin
