package usecase

import (
	"github.com/danisbagus/golang-grpc-mongodb/server/repo"
)

type IArticleUsecase interface {
	CreateArticle(data *repo.Article) (*repo.Article, error)
	GetDetail(articleID string) (*repo.Article, error)
	UpdateArticle(articleID string, data *repo.Article) (*repo.Article, error)
}

type ArticleUsecase struct {
	repo repo.IArticleRepo
}

func NewArticleUsecase(repo repo.IArticleRepo) IArticleUsecase {
	return &ArticleUsecase{
		repo: repo,
	}
}

func (r ArticleUsecase) CreateArticle(data *repo.Article) (*repo.Article, error) {
	res, err := r.repo.Create(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r ArticleUsecase) GetDetail(articleID string) (*repo.Article, error) {
	data, err := r.repo.GetOneByID(articleID)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r ArticleUsecase) UpdateArticle(articleID string, data *repo.Article) (*repo.Article, error) {
	data, err := r.repo.Update(articleID, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
