package usecase

import "github.com/danisbagus/golang-grpc-mongodb/server/repo"

type IArticleUsecase interface {
}

type ArticleUsecase struct {
	repo repo.IArticleRepo
}

func NewArticleUsecase(repo repo.IArticleRepo) IArticleUsecase {
	return &ArticleUsecase{
		repo: repo,
	}
}
