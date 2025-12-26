package repositories

import ArticleModel "Blog/internal/modules/article/models"

type ArticleRepositoryInterface interface {
	List(limit int) []ArticleModel.Article
	Find(id int) (ArticleModel.Article, error)
}
