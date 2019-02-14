.PHONY: build clean mocks report test images cmdtest

build:
	go build -o bin/jc ./cmd/jc/main.go 

clean:
	rm coverage.out
	rm -rf bin/

mocks:
	go generate ./...

report:
	go tool cover -html=coverage.out

test:
	go test -v -coverprofile=coverage.out ./...

images:
	docker build --target=cmdtest -t jc/cmdtest:latest .

cmdtest:
	docker run --rm -it jc/cmdtest
