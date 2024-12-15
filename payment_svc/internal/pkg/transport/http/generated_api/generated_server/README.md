# Go API Server for openapi

API для обработки платежей за комнаты.

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 1.0.0
- Build date: 2024-12-15T17:24:16.200317+03:00[Europe/Moscow]
- Generator version: 7.10.0


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

The server will be available on `http://localhost:8080`.

To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once image is built use
```
docker run --rm -it openapi
```