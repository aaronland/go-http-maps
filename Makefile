GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
CWD=$(shell pwd)

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/server cmd/server/main.go
