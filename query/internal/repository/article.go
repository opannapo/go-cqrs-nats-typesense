package repository

import (
	"context"
	"gcnt/internal/model"
)

var ArticleRepositoryInstance *ArticleRepositoryImpl

func InitArticleRepositoryInstance() {
	if ArticleRepositoryInstance == nil {
		ArticleRepositoryInstance = &ArticleRepositoryImpl{}
	}
}

type IArticleRepository interface {
	Get(ctx context.Context) (err error)
}

type ArticleRepositoryImpl struct {
}

func (a ArticleRepositoryImpl) Get(ctx context.Context) (article model.Article, err error) {
	return
}
