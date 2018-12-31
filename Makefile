pkgs = $(shell go list ./...)

lint:
	golint -set_exit_status ./...

test:
	go test -v -race -cover $(pkgs)