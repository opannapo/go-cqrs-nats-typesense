package service

import (
	"context"
	"gcnt/internal/repository"
	"gcnt/internal/schema"
	"github.com/rs/zerolog/log"
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
}

type articleService struct {
}

func (a *articleService) GetByID(ctx context.Context, id string) (res interface{}, err error) {
	client := repository.DbInstance.TypeSense

	res, err = client.Collection("articles").Document(id).Retrieve()
	if err != nil {
		return schema.GetResponse{}, err
	}
	if err != nil {
		log.Err(err).Caller()
		return
	}

	log.Info().Msgf("getby id retrieve : %+v", res)
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

func (a *articleService) Get(reqFilter map[string]string, ctx context.Context) (res schema.GetResponse, err error) {
	return res, err
}
