package article

// IArticle is the interface to interact with artice entity
type IArticle interface {
	Create(request *ArticleDto) (*ArticleDto, error)
	GetByID(id int64) (*ArticleDto, error)
	GetAll() ([]*ArticleDto, error)
}
