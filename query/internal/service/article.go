package service

import (
	"context"
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
}

type articleService struct {
}

func (a *articleService) Get(reqFilter map[string]string, ctx context.Context) (res schema.GetResponse, err error) {

	return res, err
}
