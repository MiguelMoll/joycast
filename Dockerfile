FROM golang:1.11.5 as builder

# User to pack templates into binary
RUN go get -u github.com/gobuffalo/packr/packr

ADD . /app
WORKDIR /app
RUN packr build -o bin/web ./frontend 

FROM debian:stretch-slim as runner
COPY --from=builder /app/bin/ /app/bin/
CMD ["bin/web"]
