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

	mysql := repository.DbInstance.Mysql
	mysql.Begin()
	article, err := repository.ArticleRepositoryInstance.Create(ctx, &newArt, mysql)
	if err != nil {
		log.Err(err).Caller()
		mysql.Rollback()
		return
	}

	err = publisher.Nats.Publish("article.created", article)
	if err != nil {
		log.Err(err).Caller().Send()
		mysql.Rollback()
		return
	}

	res = schema.CreateResponse{
		ID:    article.Id,
		Title: article.Title,
	}

	mysql.Commit()

	return res, err
}
