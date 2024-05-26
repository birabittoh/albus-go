# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

WORKDIR /build

# Download Go modules
RUN apk update && \
    apk add --no-cache \
    git

COPY . .
RUN git submodule update --init --recursive
RUN go mod download

# Build
RUN CGO_ENABLED=0 go build -trimpath -o /dist/app

# Test
FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM alpine:latest AS build-release-stage

RUN apk update && \
    apk add --no-cache \
    pandoc \
    tectonic

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /dist /app

WORKDIR /app

ENTRYPOINT ["/app/app"]
