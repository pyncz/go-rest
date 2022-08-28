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
COPY base ./base
COPY middlewares ./middlewares
COPY models ./models
COPY utils ./utils
COPY main.go ./

# Copy docs' install-n-build script
COPY ./scripts/swagger.sh ./scripts/swagger.sh

# Build docs
RUN	sh scripts/swagger.sh

# Build binary
RUN go build -o /go-rest


##
## Build final image
##
FROM alpine:latest AS build

ARG EXPOSE_PORT

WORKDIR /

# Copy public assets
COPY public /public

# Copy DB init script
COPY scripts/seed.sh /docker-entrypoint-initdb.d/

COPY --from=artifacts /go-rest /go-rest

EXPOSE $EXPOSE_PORT

ENTRYPOINT [ "/go-rest" ]
