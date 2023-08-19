package handler

import (
	"gcnt/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ArticleHandlerImpl interface {
	CreateArticle(c *fiber.Ctx) error
}

func NewArticleHandler() ArticleHandlerImpl {
	return ArticleHandler{}
}

type ArticleHandler struct{}

func (a ArticleHandler) CreateArticle(c *fiber.Ctx) error {
	s := service.ArticleServiceInstance
	err := s.Create(c)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, "1", err.Error())
	}

	return ResponseOk(c, "")
}
