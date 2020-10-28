package article

// ArticleEntity refers to the structure that is going to be stored in the DB
type ArticleEntity struct {
	ID      int64  `gorm:"column:id;type:bigserial;primary_key"`
	Title   string `gorm:"column:title;type:varchar(100)"`
	Content string `gorm:"column:content;type:text"`
	Author  string `gorm:"column:author;type:varchar(100)"`
}
