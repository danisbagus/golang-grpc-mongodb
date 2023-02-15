package handler

type ArticleResponse struct {
	Id       string `json:"id"`
	AuthorId string `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
