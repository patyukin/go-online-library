FROM golang:1.22.1-alpine3.19 AS build-env

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . /app
WORKDIR /app

RUN go mod download
RUN go mod tidy
RUN go build -gcflags="all=-N -l" -o bin/server cmd/onlib/main.go

FROM alpine:3.18

EXPOSE 8087 40001

WORKDIR /root

COPY --from=build-env /go/bin/dlv .
COPY --from=build-env /app/bin/server ./server

CMD ["/root/dlv", "--listen=:40001", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/root/server"]
