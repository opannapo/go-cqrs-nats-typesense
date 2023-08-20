package service

import (
	"context"
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
	GetByID(ctx context.Context, id string) (res interface{}, err error)
	Upsert(ctx context.Context, article schema.MessageConsume) (err error)
	Search(ctx context.Context, query, author string) (res interface{}, err error)
}

type articleService struct {
}

func (a *articleService) Search(ctx context.Context, query, author string) (res interface{}, err error) {
	searchResult, err := repository.ArticleRepositoryInstance.Search(ctx, query, author)
	if err != nil {
		log.Err(err).Caller()
		return
	}

	resHits := searchResult.Hits
	var tmpResMapping []*map[string]interface{}

	for _, hit := range *resHits {
		tmpResMapping = append(tmpResMapping, hit.Document)
	}

	res = tmpResMapping
	log.Info().Msgf("search searchResult : %+v", res)
	return
}

func (a *articleService) GetByID(ctx context.Context, id string) (res interface{}, err error) {
	res, err = repository.ArticleRepositoryInstance.GetByID(ctx, id)
	if err != nil {
		log.Err(err).Caller()
		return
	}

	log.Info().Msgf("getby id result : %+v", res)
	return
}

func (a *articleService) Upsert(ctx context.Context, article schema.MessageConsume) (err error) {
	upsert, err := repository.ArticleRepositoryInstance.Upsert(ctx, article)
	if err != nil {
		log.Err(err).Caller()
		return
	}

	log.Info().Msgf("upsert : %+v", upsert)
	return
}
