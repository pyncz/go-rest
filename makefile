##
## GO commands
##

i: # Install deps
	go mod download

dev: # Run main service
	go run .

path: # Add go to the path
	export PATH=$(go env GOPATH)/bin:$PATH

bin: # Build binaries
	go build -o /go-rest

swagi: # Install
	go install github.com/swaggo/swag/cmd/swag@latest

swagg: # Generate
	swag init


##
## Docker commands
##

build: # build containers from images
	docker compose build

start: # Start services
	docker compose start

stop: # Stop services
	docker compose stop

refresh: # Pull image, recreate and start containers
	sh scripts/refresh.sh

up: # Create and start containers
	docker compose up

deamon: # Create and start containers in background
	docker compose up -d

down: # Stop and remove containers
	docker compose down
