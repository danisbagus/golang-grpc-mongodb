package handler

import (
	"context"
	"fmt"

	"github.com/danisbagus/golang-grpc-mongodb/common/model"
	"github.com/danisbagus/golang-grpc-mongodb/server/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerHandler struct {
	usecase usecase.IArticleUsecase
}

func NewArticleHandler(usecase usecase.IArticleUsecase) *ServerHandler {
	return &ServerHandler{
		usecase: usecase,
	}
}

func (r ServerHandler) ReadArticle(ctx context.Context, req *model.ReadArticleRequest) (*model.ReadArticleResponse, error) {
	fmt.Println("Read article request")

	articleID := req.GetArticleId()
	oid, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID: %v", err),
		)
	}

	data, err := r.usecase.GetDetail(oid)
	if err != nil {
		return nil, err
	}

	return &model.ReadArticleResponse{
		Article: data,
	}, nil

}
