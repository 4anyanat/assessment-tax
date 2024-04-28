FROM golang:1.22.2 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /out/myapp

FROM debian:buster-slim

COPY --from=build /out/myapp /usr/local/bin/myapp

ENTRYPOINT ["/usr/local/bin/myapp"]
