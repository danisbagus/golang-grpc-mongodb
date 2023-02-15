package handler

type CreateArticleRequest struct {
	AuthorId string `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
