package repo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Article struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

type IArticleRepo interface {
	GetOneByID(articleID primitive.ObjectID) (*Article, error)
}

type ArticleRepo struct {
	db *mongo.Client
}

func NewArticleRepo(db *mongo.Client) IArticleRepo {
	return &ArticleRepo{
		db: db,
	}
}

func (r ArticleRepo) GetOneByID(articleID primitive.ObjectID) (*Article, error) {

	filter := bson.M{"_id": articleID}
	data := &Article{}

	collection := r.db.Database("golang_grpc_mongodb").Collection("articles")

	res := collection.FindOne(context.TODO(), filter)
	if err := res.Decode(data); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("[ReadArticle] Cannot find article with specified ID: %v, error: %v", articleID, err),
			)
		} else {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("[ReadArticle] Internal error: %v", err),
			)
		}
	}

	return data, nil
}
