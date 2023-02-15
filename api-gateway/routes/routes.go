package routes

import (
	"log"

	"github.com/danisbagus/golang-grpc-mongodb/api-gateway/handler"
	"github.com/danisbagus/golang-grpc-mongodb/common/config"
	"github.com/danisbagus/golang-grpc-mongodb/common/model"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func ApiRoutes(route *echo.Echo) {
	opts := grpc.WithInsecure()
	connection, err := grpc.Dial(config.SERVER_ARTICLE_PORT, opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	articleClient := model.NewArticleServiceClient(connection)
	articleHandler := handler.NewArticleHandler(articleClient)

	articleRoutes := route.Group(("/api/article"))
	articleRoutes.GET("", articleHandler.ListArticle)
	articleRoutes.GET("/:id", articleHandler.ReadArticle)
	articleRoutes.POST("", articleHandler.CreateArticle)
	articleRoutes.PUT("/:id", articleHandler.UpdateArticle)
	articleRoutes.DELETE("/:id", articleHandler.DeleteArticle)
}
