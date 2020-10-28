package article

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Articles implements the IArticle interface
type Articles struct {
	DBClient *gorm.DB
}

// New - creates a new instance of Articles Impl
func New(dbClient *gorm.DB) *Articles {
	article := Articles{
		DBClient: dbClient,
	}
	// Automigrate DB schema
	dbClient.AutoMigrate(&ArticleEntity{})
	return &article
}

// Create - Save an article in DB
func (inst *Articles) Create(request *ArticleDto) (*ArticleDto, error) {

	err := inst.DBClient.Transaction(func(tx *gorm.DB) error {

		newArticle := ArticleEntity{
			Title:   request.Title,
			Author:  request.Author,
			Content: request.Content,
		}
		response := tx.Create(&newArticle)
		if response.Error != nil {
			log.Println("Failed to create new article record", response.Error.Error())
			return response.Error
		}
		request.ID = newArticle.ID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return request, nil
}

// GetByID - Get article By id
func (inst *Articles) GetByID(id int64) (*ArticleDto, error) {
	articleEntity := ArticleEntity{}
	response := inst.DBClient.Where("id = ?", id).First(&articleEntity)
	if response.Error != nil {
		log.Println("No article found for id: ", id)
		return nil, response.Error
	}
	articleDto := ArticleDto{
		ID:      articleEntity.ID,
		Title:   articleEntity.Title,
		Author:  articleEntity.Author,
		Content: articleEntity.Content,
	}
	return &articleDto, nil
}

// GetAll - Get all articles from DB
func (inst *Articles) GetAll() ([]*ArticleDto, error) {

	articleEntities := []ArticleEntity{}
	response := inst.DBClient.Find(&articleEntities)
	if response.Error != nil {
		log.Println("Error while fetching articles")
		return nil, response.Error
	}

	if len(articleEntities) == 0 {
		log.Println("No articles found")
		return nil, nil
	}

	articleList := []*ArticleDto{}
	for _, ad := range articleEntities {
		articleDto := ArticleDto{
			ID:      ad.ID,
			Title:   ad.Title,
			Content: ad.Content,
			Author:  ad.Author,
		}
		articleList = append(articleList, &articleDto)
	}
	return articleList, nil
}
