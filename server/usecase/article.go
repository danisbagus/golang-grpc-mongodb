package usecase

import (
	"github.com/danisbagus/golang-grpc-mongodb/common/model"
	"github.com/danisbagus/golang-grpc-mongodb/server/dto"
	"github.com/danisbagus/golang-grpc-mongodb/server/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IArticleUsecase interface {
	GetDetail(articleID primitive.ObjectID) (*model.Article, error)
}

type ArticleUsecase struct {
	repo repo.IArticleRepo
}

func NewArticleUsecase(repo repo.IArticleRepo) IArticleUsecase {
	return &ArticleUsecase{
		repo: repo,
	}
}

func (r ArticleUsecase) GetDetail(transactionID primitive.ObjectID) (*model.Article, error) {
	data, err := r.repo.GetOneByID(transactionID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailArticleResponse(data)
	return response, nil
}
