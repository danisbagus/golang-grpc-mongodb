package dto

import (
	"github.com/danisbagus/golang-grpc-mongodb/common/model"
	"github.com/danisbagus/golang-grpc-mongodb/server/repo"
)

func NewGetDetailArticleResponse(data *repo.Article) *model.Article {
	return &model.Article{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Content:  data.Content,
		Title:    data.Title,
	}
}
