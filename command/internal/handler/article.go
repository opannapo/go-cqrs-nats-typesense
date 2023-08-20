package handler

import (
	"encoding/json"
	"gcnt/internal/schema"
	"gcnt/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
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
	ctx := c.Context()
	body := c.Body()

	payload := &schema.CreateRequest{}
	if err := json.Unmarshal(body, payload); err != nil {
		log.Err(err).Caller().Send()
		return ResponseError(c, http.StatusBadRequest, "1", err.Error())
	}

	var validate = validator.New()
	err := validate.Struct(payload)
	if err != nil {
		log.Err(err).Caller().Send()
		return ResponseError(c, http.StatusBadRequest, "2", err.Error())
	}

	res, err := service.ArticleServiceInstance.Create(*payload, ctx)
	if err != nil {
		log.Err(err).Caller().Send()
		return ResponseError(c, http.StatusInternalServerError, "2", err.Error())
	}

	return ResponseOk(c, res)
}
