# Setup GO path
export PATH=$(go env GOPATH)/bin:$PATH

# Install swagger
go install github.com/swaggo/swag/cmd/swag@latest

# Build docs
swag init
