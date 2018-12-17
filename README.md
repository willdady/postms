# PostMS

PostMS is a Go microservice exposing a RESTful API for managing blog-like posts and comments in a PostgreSQL database.

Note this is built with the intention of being used as an internal microservice and should only be exposed via an API Gateway (or GraphQL) layer.

## Requirements

PostMS requires access to a Postgres database. Connection details must be provided as environment variables. The default values are as follows:

```
PG_HOST=0.0.0.0
PG_PORT=5432
PG_USER=postgres
PG_DB=postgres
PG_PASSWORD=mysecretpassword
PG_SSL_MODE=disable
PORT=8080
```

In production you should also disable [gin's](https://github.com/gin-gonic/gin) debug logging:

```
GIN_MODE=production
```

## Development

Run a Postgres database easily with Docker:

```
docker run --rm -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 postgres:11.1
```

Run the app:

```
go run cmd/postms/postms.go
```
