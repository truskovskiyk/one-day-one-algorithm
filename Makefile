pkgs = $(shell go list ./...)

lint:
	golint ./...

test:
	go test -v -race -cover $(pkgs)