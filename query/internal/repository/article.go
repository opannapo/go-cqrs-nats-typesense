package repository

import (
	"context"
	"gcnt/internal/schema"
	"github.com/rs/zerolog/log"
	"github.com/typesense/typesense-go/typesense/api"
)

var ArticleRepositoryInstance IArticleRepository

func InitArticleRepositoryInstance() {
	if ArticleRepositoryInstance == nil {
		ArticleRepositoryInstance = &articleRepositoryImpl{}
	}
}

type IArticleRepository interface {
	GetByID(ctx context.Context, id string) (res map[string]interface{}, err error)
	Upsert(ctx context.Context, article schema.MessageConsume) (upsert map[string]interface{}, err error)
	Search(ctx context.Context, query, author string) (searchResult *api.SearchResult, err error)
}

type articleRepositoryImpl struct {
}

func (a articleRepositoryImpl) GetByID(ctx context.Context, id string) (res map[string]interface{}, err error) {
	client := DbInstance.TypeSense

	res, err = client.Collection("articles").Document(id).Retrieve()
	if err != nil {
		log.Err(err).Caller().Send()
		return
	}

	return
}

func (a articleRepositoryImpl) Upsert(ctx context.Context, article schema.MessageConsume) (upsert map[string]interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a articleRepositoryImpl) Search(ctx context.Context, query, author string) (searchResult *api.SearchResult, err error) {
	//TODO implement me
	panic("implement me")
}
