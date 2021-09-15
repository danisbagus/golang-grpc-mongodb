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

	// createArticle(articleClient)
	// readArticle(articleClient)
	// updateArticle(articleClient)
	// readArticle(articleClient)
	deleteArticle(articleClient)
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

func updateArticle(client model.ArticleServiceClient) {
	fmt.Println("Update article")

	req := &model.UpdateArticleRequest{
		Article: &model.Article{
			Id:       "613d7b91a5253d1732be46f5",
			AuthorId: "3",
			Title:    "MSI",
			Content:  "Visual Studio Code is on of the best laptop brand in the world",
		},
	}

	res, err := client.UpdateArticle(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling UpdateArticle RPC: %v\n", err)
	}

	log.Printf("Article has been updated: %v\n", res)
}

func deleteArticle(client model.ArticleServiceClient) {
	fmt.Println("Delete article")

	req := &model.DeleteArticleRequest{
		ArticleId: "61413edfbeae4038836d9e0e",
	}

	res, err := client.DeleteArticle(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling DeleteArticle RPC: %v\n", err)
	}

	log.Printf("Article has been deleted: %v\n", res)
}
