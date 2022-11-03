.PHONY: test cover

test:
	go test -coverprofile=cover.out ./...

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=./cover.out -o cover.html
	open cover.html
