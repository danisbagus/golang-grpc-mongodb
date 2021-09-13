package main

import (
	"context"
	"fmt"
	"log"

	"github.com/danisbagus/golang-grpc-mongodb/common/config"
	"github.com/danisbagus/golang-grpc-mongodb/common/model"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Running Client...")

	opts := grpc.WithInsecure()
	connection, err := grpc.Dial(config.SERVER_ARTICLE_PORT, opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer connection.Close()
	articleClient := model.NewArticleServiceClient(connection)

	createArticle(articleClient)
	// readArticle(articleClient)
}

func readArticle(client model.ArticleServiceClient) {
	fmt.Println("Reading the article")

	req := &model.ReadArticleRequest{ArticleId: "613d7b91a5253d1732be46f5"}

	res, err := client.ReadArticle(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling ReadArticle RPC: %v\n", err)
	}
	log.Printf("Article was read: %v\n", res)
}

func createArticle(client model.ArticleServiceClient) {
	fmt.Println("Creating new article")

	req := &model.CreateArticleRequest{
		Article: &model.Article{
			AuthorId: "3",
			Title:    "Editor IDE",
			Content:  "Visual Studio Code is on of the best Editor IDE in the world",
		},
	}

	res, err := client.CreateArticle(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling CreateArticle RPC: %v\n", err)
	}
	log.Printf("Article has been created: %v\n", res)
}
