package services

import (
	"Blog/internal/modules/article/requests/articles"
	ArticleResponse "Blog/internal/modules/article/responses"
	UserResponse "Blog/internal/modules/user/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
	Find(id int) (ArticleResponse.Article, error)
	StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error)
}
