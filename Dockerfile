# Go version 1.19.1
FROM golang:alpine3.16 as builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main \
    -ldflags '-s -w -X main.mode=prod'

FROM scratch as runner

FROM scratch
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/main /app/main

COPY ./templates /templates/

ENTRYPOINT ["/app/main"]
