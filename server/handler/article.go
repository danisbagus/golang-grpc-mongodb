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

	article := dto.NewCreateArticleRequest(req.GetArticle())

	data, err := r.usecase.CreateArticle(article)
	if err != nil {
		return nil, err
	}

	response := dto.NewCreateArticleResponse(data)

	return response, nil
}

func (r ServerHandler) LisArticle(_ *model.LisArticleRequest, stream model.ArticleService_LisArticleServer) error {
	fmt.Println("List Article request")

	data, err := r.usecase.GetAll()
	if err != nil {
		return err
	}

	response := dto.NewGetListArticleResponse(data)

	for _, v := range response {
		stream.Send(&v)
	}
	return nil
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

func (r ServerHandler) UpdateArticle(ctx context.Context, req *model.UpdateArticleRequest) (*model.UpdateArticleResponse, error) {
	fmt.Println("Update article request")

	article := dto.NewUpdateArticleRequest(req.GetArticle())
	articleID := req.GetArticle().GetId()

	data, err := r.usecase.UpdateArticle(articleID, article)
	if err != nil {
		return nil, err
	}

	response := dto.NewUpdateArticleResponse(data)

	return response, nil
}

func (r ServerHandler) DeleteArticle(ctx context.Context, req *model.DeleteArticleRequest) (*model.DeleteArticleResponse, error) {
	fmt.Println("Delete article request")

	articleID := req.GetArticleId()

	err := r.usecase.DeleteArticle(articleID)
	if err != nil {
		return nil, err
	}

	return &model.DeleteArticleResponse{
		ArticleId: req.ArticleId,
	}, nil
}
