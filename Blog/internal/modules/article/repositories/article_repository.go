package repositories

import (
	ArticleModel "Blog/internal/modules/article/models"
	"Blog/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{DB: database.Connection()}
}

func (articleRepository *ArticleRepository) List(limit int) []ArticleModel.Article {
	var articles []ArticleModel.Article
	articleRepository.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}

func (articleRepository *ArticleRepository) Find(id int) (ArticleModel.Article, error) {
	var article ArticleModel.Article
	if err := articleRepository.DB.Joins("User").First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}
