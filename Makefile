.PHONY: build

build:
	go build -o build/chrono cmd/chrono/main.go

test:
	go test ./... -v

cover:
	go test ./... -coverpkg=./... -coverprofile cp.out
	go tool cover -html=cp.out
