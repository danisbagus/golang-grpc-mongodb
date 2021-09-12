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

	readArticle(articleClient)
}

func readArticle(client model.ArticleServiceClient) {
	fmt.Println("Reading the article")

	req := &model.ReadArticleRequest{ArticleId: "613d7b91a5253d1732be46f5"}

	res, err := client.ReadArticle(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v\n", err)
	}
	log.Printf("Article was read: %v\n", res)
}
