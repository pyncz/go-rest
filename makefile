##
## GO commands
##

i: # Install deps
	go mod download

dev: # Run main service
	go run .

bin: # Build binaries
	go build -o .output/bin


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
	sh ./scripts/refresh.sh

up: # Create and start containers
	docker compose up

deamon: # Create and start containers in background
	docker compose up -d

down: # Stop and remove containers
	docker compose down
