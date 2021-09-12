package handler

import "github.com/danisbagus/golang-grpc-mongodb/server/usecase"

type ServerHandler struct {
	usecase usecase.IArticleUsecase
}

func NewArticleHandler(usecase usecase.IArticleUsecase) *ServerHandler {
	return &ServerHandler{
		usecase: usecase,
	}
}
