package handler

import (
	"gcnt/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ArticleHandlerImpl interface {
	GetArticleByID(c *fiber.Ctx) error
	Search(c *fiber.Ctx) error
}

func NewArticleHandler() ArticleHandlerImpl {
	return ArticleHandler{}
}

type ArticleHandler struct{}

func (a ArticleHandler) Search(c *fiber.Ctx) error {
	ctx := c.Context()

	qQuery := c.Query("query")
	qAuthor := c.Query("author")

	res, err := service.ArticleServiceInstance.Search(ctx, qQuery, qAuthor)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, "2", err.Error())
	}

	return ResponseOk(c, res)
}

func (a ArticleHandler) GetArticleByID(c *fiber.Ctx) error {
	ctx := c.Context()

	articleID := c.Params("articleID")

	if articleID == "" {
		return ResponseError(c, http.StatusBadRequest, "3", "articleID empty")
	}

	res, err := service.ArticleServiceInstance.GetByID(ctx, articleID)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, "2", err.Error())
	}

	return ResponseOk(c, res)
}
