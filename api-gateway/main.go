package main

import (
	"github.com/danisbagus/golang-grpc-mongodb/api-gateway/routes"
	"github.com/danisbagus/golang-grpc-mongodb/common/config"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.ApiRoutes(e)

	e.Logger.Fatal(e.Start(config.SERVER_API_GATEWAY_PORT))
}
