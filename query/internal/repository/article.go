package repository

import (
	"context"
	"fmt"
	"gcnt/internal/schema"
	"github.com/rs/zerolog/log"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"strconv"
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
	client := DbInstance.TypeSense

	document :=
		struct {
			ID      string `json:"id" `
			Author  string `json:"author" `
			Title   string `json:"title" `
			Body    string `json:"body" `
			Created int64  `json:"created" `
		}{
			ID:      strconv.FormatInt(article.ID, 10),
			Author:  article.Author,
			Title:   article.Title,
			Body:    article.Body,
			Created: article.Created.Unix(),
		}

	upsert, err = client.Collection("articles").Documents().Upsert(document)
	if err != nil {
		log.Err(err).Caller().Send()
		return
	}

	return
}

func (a articleRepositoryImpl) Search(ctx context.Context, query, author string) (searchResult *api.SearchResult, err error) {
	client := DbInstance.TypeSense

	sort := "created:desc"
	searchParameters := &api.SearchCollectionParams{
		Q:        query,
		QueryBy:  "title, body",
		FilterBy: pointer.String(fmt.Sprintf("author:%s", author)),
		SortBy:   &sort,
	}

	searchResult, err = client.Collection("articles").Documents().Search(searchParameters)
	if err != nil {
		log.Err(err).Caller().Send()
		return
	}

	return
}
