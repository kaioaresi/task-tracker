tidy:
	go mod download
	go mod tidy -v

test:
	go test -v .

coverage:
	go test -cover .
	go test -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out
