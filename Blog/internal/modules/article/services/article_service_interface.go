package services

import ArticleResponse "Blog/internal/modules/article/responses"

type ArticleServiceInterface interface {
	GetFeaturedArticles() ArticleResponse.Articles
	GetStoriesArticles() ArticleResponse.Articles
}
