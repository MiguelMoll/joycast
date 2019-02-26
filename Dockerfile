FROM golang:1.11.5 as builder

# USed to pack templates into the binary
RUN go get -u github.com/gobuffalo/packr/packr

ADD . /app
WORKDIR /app
RUN packr build -o bin/web ./frontend 

FROM debian:stretch-slim as runner
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /app
COPY --from=builder /app /app
CMD ["bin/web"]
