## Go REST
My lil playground to try GO.

![GO](.github/banner.jpg)


# Development

## Let's GO

### 📦 Install dependencies
```sh
make i
```

### 🔧 Launch dev server
```sh
make dev
```

> **Warning**
> Don't forget to set environment variables via `.env` file in the root dir. See `.env.template` to know what vars are expected.


### 📖 Generate swagger docs
Run this to install swagger generator:
```sh
go get -u github.com/swaggo/swag/cmd/swag
```

To generate docs, run:
```sh
make swagi
```

> **NOTE**
> Make sure you've added GOPATH to your PATH env variable
> ```
> export PATH=$(go env GOPATH)/bin:$PATH
> ```


## Docker Commands

### 🔄 Restart containers
This command pulls web service's image, recreate and start containers. You will probably use it on server when refresh is required after an update.

```sh
make restart
```

### 🚀 Up containers
These commands will build containers by their images and create / start them.

```sh
make build
make up
```

---
For more commands, see [makefile](makefile).
