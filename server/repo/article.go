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
	Create(data *Article) (*Article, error)
	GetOneByID(articleID string) (*Article, error)
	Update(articleID string, data *Article) (*Article, error)
	Delete(articleID string) error
}

type ArticleRepo struct {
	db *mongo.Client
}

func NewArticleRepo(db *mongo.Client) IArticleRepo {
	return &ArticleRepo{
		db: db,
	}
}

func (r ArticleRepo) Create(data *Article) (*Article, error) {
	collection := r.db.Database("golang_grpc_mongodb").Collection("articles")

	res, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[Create] Error on create article: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[Create] Cannot convert to OID: %v", err),
		)
	}

	data.ID = oid
	return data, nil
}

func (r ArticleRepo) GetOneByID(articleID string) (*Article, error) {
	oid, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID: %v", err),
		)
	}

	filter := bson.M{"_id": oid}
	data := &Article{}

	collection := r.db.Database("golang_grpc_mongodb").Collection("articles")

	res := collection.FindOne(context.TODO(), filter)
	if err := res.Decode(data); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("[GetOneByID] Cannot find article with specified ID: %v, error: %v", oid, err),
			)
		} else {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("[GetOneByID] Internal error: %v", err),
			)
		}
	}

	return data, nil
}

func (r ArticleRepo) Update(articleID string, data *Article) (*Article, error) {
	oid, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID: %v", err),
		)
	}

	filter := bson.M{"_id": oid}
	article := &Article{}

	collection := r.db.Database("golang_grpc_mongodb").Collection("articles")

	res := collection.FindOne(context.TODO(), filter)
	if err := res.Decode(article); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("[Update] Cannot find article with specified ID: %v, error: %v", oid, err),
			)
		} else {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("[Update] Internal error: %v", err),
			)
		}
	}

	article.AuthorID = data.AuthorID
	article.Title = data.Title
	article.Content = data.Content

	_, updatErr := collection.ReplaceOne(context.TODO(), filter, article)
	if updatErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("[Update] Error on update article: %v", err),
		)
	}

	return article, nil
}

func (r ArticleRepo) Delete(articleID string) error {
	oid, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		return status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID: %v", err),
		)
	}

	filter := bson.M{"_id": oid}
	article := &Article{}

	collection := r.db.Database("golang_grpc_mongodb").Collection("articles")

	res := collection.FindOne(context.TODO(), filter)
	if err := res.Decode(article); err != nil {
		if err == mongo.ErrNoDocuments {
			return status.Errorf(
				codes.NotFound,
				fmt.Sprintf("[Delete] Cannot find article with specified ID: %v, error: %v", oid, err),
			)
		} else {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("[Delete] Internal error: %v", err),
			)
		}
	}

	_, errDelete := collection.DeleteOne(context.TODO(), filter)
	if errDelete != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("[Delete] Error on delete article: %v", err),
		)
	}

	return nil
}
