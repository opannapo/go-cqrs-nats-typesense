package handler

import "github.com/gofiber/fiber/v2"

type ArticleHandlerImpl interface {
	CreateArticle(c *fiber.Ctx) (err error)
}

func NewArticleHandler() ArticleHandlerImpl {
	return ArticleHandler{}
}

type ArticleHandler struct{}

func (a ArticleHandler) CreateArticle(c *fiber.Ctx) (err error) {
	//TODO implement me
	panic("implement me")
}
