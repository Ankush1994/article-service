package article

// ArticleDto represents an Article API request and response
type ArticleDto struct {
	ID      int64  `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
}
