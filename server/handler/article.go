package handler

import (
	"context"
	"fmt"

	"github.com/danisbagus/golang-grpc-mongodb/common/model"
	"github.com/danisbagus/golang-grpc-mongodb/server/dto"
	"github.com/danisbagus/golang-grpc-mongodb/server/usecase"
)

type ServerHandler struct {
	usecase usecase.IArticleUsecase
}

func NewArticleHandler(usecase usecase.IArticleUsecase) *ServerHandler {
	return &ServerHandler{
		usecase: usecase,
	}
}

func (r ServerHandler) CreateArticle(ctx context.Context, req *model.CreateArticleRequest) (*model.CreateArticleResponse, error) {
	fmt.Println("Create article request")

	article := dto.NewCreateArticleRequest(req.Article)

	data, err := r.usecase.CreateArticle(article)
	if err != nil {
		return nil, err
	}

	response := dto.NewCreateArticleResponse(data)

	return response, nil
}

func (r ServerHandler) ReadArticle(ctx context.Context, req *model.ReadArticleRequest) (*model.ReadArticleResponse, error) {
	fmt.Println("Read article request")

	articleID := req.GetArticleId()

	data, err := r.usecase.GetDetail(articleID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailArticleResponse(data)

	return response, nil

}
