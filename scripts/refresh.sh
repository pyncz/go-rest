# Update image
docker compose pull web

# Restart containers
docker compose stop
docker compose up -d
