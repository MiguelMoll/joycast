.PHONY: build clean mocks report test images cmdtest

build:
	go build -o bin/web ./cmd/web/main.go 
	go build -o bin/minion ./cmd/minion/main.go

clean:
	rm -f coverage.out
	rm -rf bin/

mocks:
	go generate ./...

report: test
	go tool cover -html=coverage.out

test:
	go test -v -coverprofile=coverage.out ./...

images:
	docker build -t web -f Dockerfile .
