FROM golang:1.11.5 as builder

# User to pack templates into binary
RUN go get -u github.com/gobuffalo/packr/packr

ADD . /app
WORKDIR /app
RUN packr build -o bin/web ./frontend 

FROM debian:stretch-slim as runner
WORKDIR /app
COPY --from=builder /app /app
COPY --from=builder /app/frontend/client_secret.json /app/bin
CMD ["bin/web"]
