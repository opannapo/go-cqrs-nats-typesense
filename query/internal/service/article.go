package service

import (
	"context"
	"fmt"
	"gcnt/internal/repository"
	"gcnt/internal/schema"
	"github.com/rs/zerolog/log"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
	"strconv"
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
	client := repository.DbInstance.TypeSense

	sort := "created:desc"
	searchParameters := &api.SearchCollectionParams{
		Q:        query,
		QueryBy:  "title, body",
		FilterBy: pointer.String(fmt.Sprintf("author:%s", author)),
		SortBy:   &sort,
	}

	search, err := client.Collection("articles").Documents().Search(searchParameters)
	if err != nil {
		log.Err(err).Caller().Send()
		return schema.GetResponse{}, err
	}

	resHits := search.Hits
	var tmpResMapping []*map[string]interface{}

	for _, hit := range *resHits {
		tmpResMapping = append(tmpResMapping, hit.Document)
	}

	res = tmpResMapping
	log.Info().Msgf("getby id retrieve : %+v", res)
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
	client := repository.DbInstance.TypeSense

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

	upsert, err := client.Collection("articles").Documents().Upsert(document)
	if err != nil {
		log.Err(err).Caller()
		return
	}

	log.Info().Msgf("upsert : %+v", upsert)
	return
}
