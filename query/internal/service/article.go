package service

import (
	"context"
	"gcnt/internal/model"
	"gcnt/internal/schema"
)

var ArticleServiceInstance IArticleService

func InitArticleServiceInstance() {
	if ArticleServiceInstance == nil {
		ArticleServiceInstance = &articleService{}
	}
}

type IArticleService interface {
	Get(reqFilter map[string]string, ctx context.Context) (res schema.GetResponse, err error)
	Create(article model.Article, ctx context.Context) (res schema.GetResponse, err error)
}

type articleService struct {
}

func (a *articleService) Create(article model.Article, ctx context.Context) (res schema.GetResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleService) Get(reqFilter map[string]string, ctx context.Context) (res schema.GetResponse, err error) {

	return res, err
}
