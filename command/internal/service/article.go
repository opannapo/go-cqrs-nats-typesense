package service

import (
	"context"
	"gcnt/internal/model"
	"gcnt/internal/repository"
	"gcnt/internal/schema"
	"gcnt/internal/stream/publisher"
	"github.com/rs/zerolog/log"
	"time"
)

var ArticleServiceInstance IArticleService

func InitArticleServiceInstance() {
	if ArticleServiceInstance == nil {
		ArticleServiceInstance = &articleService{}
	}
}

type IArticleService interface {
	Create(req schema.CreateRequest, ctx context.Context) (res schema.CreateResponse, err error)
}

type articleService struct {
}

func (a *articleService) Create(req schema.CreateRequest, ctx context.Context) (res schema.CreateResponse, err error) {
	newArt := model.Article{
		Author:  req.Author,
		Title:   req.Title,
		Body:    req.Body,
		Created: time.Now(),
	}

	trx := repository.DbInstance.Mysql.Begin()
	article, err := repository.ArticleRepositoryInstance.Create(ctx, &newArt, trx)
	if err != nil {
		log.Err(err).Caller()
		trx.Rollback()
		return
	}

	err = publisher.Nats.Publish("article.created", article)
	if err != nil {
		log.Err(err).Caller().Send()
		trx.Rollback()
		return
	}

	res = schema.CreateResponse{
		ID:    article.Id,
		Title: article.Title,
	}

	trx.Commit()

	return res, err
}
