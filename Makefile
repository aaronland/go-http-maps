GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
CWD=$(shell pwd)

example:
	go run cmd/example/main.go \
		-initial-view '-122.384292,37.621131,13'
