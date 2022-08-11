# syntax=docker/dockerfile:1


##
## Build artifacts
##
FROM golang:1.19-alpine AS artifacts

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy sources
COPY api ./api
COPY utils ./utils
COPY models ./models
COPY main.go ./

# Copy DB init script
COPY ./scripts/seed.sh /scripts/seed.sh

# Build binary
RUN go build -o /bin/go-rest


##
## Build docker image
##
FROM alpine:latest AS build

ARG EXPOSE_PORT

WORKDIR /

COPY --from=artifacts /bin/go-rest /bin/go-rest
COPY --from=artifacts /scripts/seed.sh /docker-entrypoint-initdb.d/

EXPOSE $EXPOSE_PORT

ENTRYPOINT [ "/bin/go-rest" ]
