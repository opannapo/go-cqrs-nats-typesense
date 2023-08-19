package service

import (
	"gcnt/internal/model"
	"gcnt/internal/repository"
	"github.com/gofiber/fiber/v2"
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
	Create(c *fiber.Ctx) (err error)
}

type articleService struct {
}

func (a *articleService) Create(c *fiber.Ctx) (err error) {
	newArt := model.Article{
		Id:      0,
		Author:  "",
		Title:   "",
		Body:    "",
		Created: time.Now(),
	}

	r := repository.ArticleRepositoryInstance
	err = r.Create(&newArt, repository.DbInstance.Mysql)
	if err != nil {
		log.Err(err).Caller()
		return err
	}

	return
}
