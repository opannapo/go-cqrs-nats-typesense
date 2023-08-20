package service

import (
	"context"
	"gcnt/internal/model"
	"gcnt/internal/repository"
	"gcnt/internal/schema"
	"github.com/rs/zerolog/log"
)

var ArticleServiceInstance IArticleService

func InitArticleServiceInstance() {
	if ArticleServiceInstance == nil {
		ArticleServiceInstance = &articleService{}
	}
}

type IArticleService interface {
	Get(reqFilter map[string]string, ctx context.Context) (res schema.GetResponse, err error)
	Upsert(article model.Article, ctx context.Context) (err error)
}

type articleService struct {
}

func (a *articleService) Upsert(article model.Article, ctx context.Context) (err error) {
	client := repository.DbInstance.TypeSense
	upsert, err := client.Collection("articles").Documents().Upsert(article)
	if err != nil {
		log.Err(err).Caller()
		return
	}

	log.Info().Msgf("upsert : %+v", upsert)
	return
}

func (a *articleService) Get(reqFilter map[string]string, ctx context.Context) (res schema.GetResponse, err error) {

	return res, err
}
