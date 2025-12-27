package services

import (
	ArticleModel "Blog/internal/modules/article/models"
	ArticleRepository "Blog/internal/modules/article/repositories"
	"Blog/internal/modules/article/requests/articles"
	ArticleResponse "Blog/internal/modules/article/responses"
	UserResponse "Blog/internal/modules/user/responses"
	"errors"
)

type ArticleService struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: ArticleRepository.New(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(4)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() ArticleResponse.Articles {
	articles := articleService.articleRepository.List(6)
	return ArticleResponse.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (ArticleResponse.Article, error) {
	var response ArticleResponse.Article
	article := articleService.articleRepository.Find(id)

	if article.ID == 0 {
		return response, errors.New("Article not found")
	}
	return ArticleResponse.ToArticle(article), nil
}

func (articleService *ArticleService) StoreAsUser(request articles.StoreRequest, user UserResponse.User) (ArticleResponse.Article, error) {
	var article ArticleModel.Article
	article.Title = request.Title
	article.Content = request.Content
	article.UserID = user.ID
	newArticle := articleService.articleRepository.Create(article)
	if newArticle.ID == 0 {
		return ArticleResponse.Article{}, errors.New("Error in creating article")
	}
	return ArticleResponse.ToArticle(newArticle), nil
}
