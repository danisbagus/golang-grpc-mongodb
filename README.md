# golang-grpc-mongodb

Simple CRUD api using Go, gRPC, and MongoDB

## Requirements

- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Docker-compose](https://docs.docker.com/compose/) for running MongoDB.

## Setup

Create MongoDB container

```bash
cd resource/docker && docker-compose up
```

## Run the service

```

Get Go packages

```bash
go get .
```

Run The server

```bash
go run server/main.go
```

Run The api gateway

```bash
go run api-gateway/main.go
```