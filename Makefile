pkgs = $(shell go list ./...)

test:
	go test -v -race -cover $(pkgs)