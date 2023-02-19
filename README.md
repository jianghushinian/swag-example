# Swagger Example API

## Getting started

1. Gen doc

```bash
$ make swag-init
```

2. Run app

```bash
$ go run cmd/main.go
```

3. Open Swagger

   Swagger UI:
    - swagger v1 [http://localhost:8080/swagger/v1/index.html](http://localhost:8080/swagger/v1/index.html)
    - swagger v2 [http://localhost:8080/swagger/v2/index.html](http://localhost:8080/swagger/v2/index.html)

   ReDoc UI:
    - redoc v1 [http://localhost:8080/redoc/v1](http://localhost:8080/redoc/v1)
    - redoc v2 [http://localhost:8080/redoc/v2](http://localhost:8080/redoc/v2)

## About the Project

This project was inspired by [swag/example](https://github.com/swaggo/swag/tree/master/example)
and [gin-swagger/example](https://github.com/swaggo/gin-swagger/tree/master/example).
