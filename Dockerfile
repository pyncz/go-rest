# syntax=docker/dockerfile:1


##
## Build artifacts
##
FROM golang:1.19-alpine AS artifacts

WORKDIR /app

# Setup GO path
RUN export PATH=$(go env GOPATH)/bin:$PATH

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Install swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy sources
COPY api ./api
COPY base ./base
COPY middlewares ./middlewares
COPY models ./models
COPY utils ./utils
COPY docs ./docs
COPY main.go ./

# Copy public assets
COPY public /public

# Copy DB init script
COPY ./scripts/seed.sh /scripts/seed.sh

# Build docs
RUN	swag init

# Build binary
RUN go build -o /go-rest


##
## Build docker image
##
FROM alpine:latest AS build

ARG EXPOSE_PORT

WORKDIR /

COPY --from=artifacts /public /public
COPY --from=artifacts /go-rest /go-rest
COPY --from=artifacts /scripts/seed.sh /docker-entrypoint-initdb.d/

EXPOSE $EXPOSE_PORT

ENTRYPOINT [ "/go-rest" ]
