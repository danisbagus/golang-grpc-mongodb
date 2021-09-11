package repo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Article struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

type IArticleRepo interface {
}

type ArticleRepo struct {
	db *mongo.Client
}

func NewArticleRepo(db *mongo.Client) IArticleRepo {
	return &ArticleRepo{
		db: db,
	}
}
