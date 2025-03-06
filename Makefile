GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
CWD=$(shell pwd)

example:
	go run cmd/example/main.go
