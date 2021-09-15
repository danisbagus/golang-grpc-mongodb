package dto

import (
	"github.com/danisbagus/golang-grpc-mongodb/common/model"
	"github.com/danisbagus/golang-grpc-mongodb/server/repo"
)

func NewGetDetailArticleResponse(data *repo.Article) *model.ReadArticleResponse {
	article := &model.Article{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}

	return &model.ReadArticleResponse{
		Article: article,
	}
}

func NewCreateArticleResponse(data *repo.Article) *model.CreateArticleResponse {
	article := &model.Article{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}

	return &model.CreateArticleResponse{
		Article: article,
	}
}

func NewCreateArticleRequest(data *model.Article) *repo.Article {
	return &repo.Article{
		AuthorID: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}

func NewUpdateArticleRequest(data *model.Article) *repo.Article {
	return &repo.Article{
		AuthorID: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}

func NewUpdateArticleResponse(data *repo.Article) *model.UpdateArticleResponse {
	article := &model.Article{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}

	return &model.UpdateArticleResponse{
		Article: article,
	}
}
