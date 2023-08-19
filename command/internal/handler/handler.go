package handler

import (
	"github.com/gofiber/fiber/v2"
)

func ResponseError(c *fiber.Ctx, httpStatus int, internalErrCode, errMsg string) (err error) {
	resErr := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    internalErrCode,
		Message: errMsg,
	}
	return c.Status(httpStatus).JSON(&resErr)
}

func ResponseOk(c *fiber.Ctx, data interface{}) (err error) {
	res := struct {
		Code string      `json:"code"`
		Data interface{} `json:"data"`
	}{
		Code: "0",
		Data: data,
	}
	return c.Status(200).JSON(&res)
}
