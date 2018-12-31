pkgs = $(shell go list ./...)

ling:
	golint ./...

test:
	go test -v -race -cover $(pkgs)