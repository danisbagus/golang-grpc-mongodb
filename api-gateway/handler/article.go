package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/danisbagus/golang-grpc-mongodb/common/model"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	client model.ArticleServiceClient
}

func NewArticleHandler(client model.ArticleServiceClient) *ArticleHandler {
	return &ArticleHandler{
		client: client,
	}
}

func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	reqBody := CreateArticleRequest{}
	if err := c.Bind(&reqBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	reqModel := &model.CreateArticleRequest{
		Article: &model.Article{
			AuthorId: reqBody.AuthorId,
			Title:    reqBody.Title,
			Content:  reqBody.Content,
		},
	}

	res, err := h.client.CreateArticle(context.Background(), reqModel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": fmt.Sprintf("error while calling CreateArticle RPC: %v", err),
			"data":    nil,
		})
	}

	data := ArticleResponse{
		Id:       res.Article.Id,
		AuthorId: res.Article.AuthorId,
		Title:    res.Article.Title,
		Content:  res.Article.Content,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully create article",
		"data":    data,
	})
}

func (h *ArticleHandler) ListArticle(c echo.Context) error {
	data := make([]ArticleResponse, 0)
	req := &model.LisArticleRequest{}
	stream, err := h.client.LisArticle(context.Background(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": fmt.Sprintf("error while calling ListArticle RPC: %v", err),
			"data":    nil,
		})
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": fmt.Sprintf("error while stream article list data: %v", err),
				"data":    nil,
			})
		}

		article := ArticleResponse{
			Id:       res.Article.Id,
			AuthorId: res.Article.AuthorId,
			Title:    res.Article.Title,
			Content:  res.Article.Content,
		}

		data = append(data, article)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully get list article",
		"data":    data,
	})
}
