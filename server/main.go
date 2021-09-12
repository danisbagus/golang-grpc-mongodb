package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/danisbagus/golang-grpc-mongodb/common/config"
	"github.com/danisbagus/golang-grpc-mongodb/common/model"
	"github.com/danisbagus/golang-grpc-mongodb/server/handler"
	"github.com/danisbagus/golang-grpc-mongodb/server/repo"
	"github.com/danisbagus/golang-grpc-mongodb/server/usecase"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
)

type server struct {
	usecase usecase.IArticleUsecase
}

func main() {
	fmt.Println("Running server...")

	client := GetClient()
	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to MongoDB", err)
	} else {
		fmt.Println("Connected to MongoDB")
	}

	listen, err := net.Listen("tcp", config.SERVER_ARTICLE_PORT)
	if err != nil {
		log.Fatalf("Failed to listen. %v", err)
	}

	opts := []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)

	articleRepo := repo.NewArticleRepo(client)
	articleUseCase := usecase.NewArticleUsecase(articleRepo)
	handler := handler.NewArticleHandler(articleUseCase)

	model.RegisterArticleServiceServer(srv, handler)

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve. %v", err)
	}
}

func GetClient() *mongo.Client {
	var cred options.Credential

	cred.AuthSource = "admin"
	cred.Username = "root"
	cred.Password = "danisbagus"

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred) // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect MongoDB %v", err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
