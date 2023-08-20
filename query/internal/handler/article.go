package handler

import (
	"github.com/gofiber/fiber/v2"
)

type ArticleHandlerImpl interface {
	GetArticle(c *fiber.Ctx) error
}

func NewArticleHandler() ArticleHandlerImpl {
	return ArticleHandler{}
}

type ArticleHandler struct{}

func (a ArticleHandler) GetArticle(c *fiber.Ctx) error {
	/*ctx := c.Context()
	body := c.Body()

	payload := &schema.CreateRequest{}
	if err := json.Unmarshal(body, payload); err != nil {
		log.Err(err).Caller()
		return ResponseError(c, http.StatusBadRequest, "1", err.Error())
	}

	var validate = validator.New()
	err := validate.Struct(payload)
	if err != nil {
		log.Print(err)
		return ResponseError(c, http.StatusBadRequest, "2", err.Error())
	}

	res, err := service.ArticleServiceInstance.Create(*payload, ctx)
	if err != nil {
		return ResponseError(c, http.StatusInternalServerError, "2", err.Error())
	}

	return ResponseOk(c, res)*/
	return ResponseOk(c, "")
}
