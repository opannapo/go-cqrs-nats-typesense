package handler_test

import (
	"bou.ke/monkey"
	article_handler "gcnt/internal/handler"
	"gcnt/internal/schema"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func (h *HandlerTestSuite) TestCreateArticle() {
	type testCase struct {
		caseName         string
		expectedLogic    func(fctx *fiber.Ctx, c testCase)
		expectedResponse func(actualRes *http.Response)
		json             string
	}
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 1, 1, 20, 34, 58, 651387237, time.UTC)
	})

	cases := []testCase{
		{
			caseName: "Success",
			json: `{
					"author": "opannapo",
					"title": "Motor KLX",
					"body": "Ini adalah artikel tentang kendaraan roda dua motor Kawasaki KLX"
				}`,
			expectedLogic: func(fctx *fiber.Ctx, c testCase) {
				ctx := fctx.Context()

				req := &schema.CreateRequest{
					Author: "opannapo",
					Title:  "Motor KLX",
					Body:   "Ini adalah artikel tentang kendaraan roda dua motor Kawasaki KLX",
				}
				res := schema.CreateResponse{
					ID:    0,
					Title: "Motor KLX",
				}
				h.mockService.EXPECT().
					Create(*req, ctx).
					Return(res, nil)
			},
			expectedResponse: func(actualRes *http.Response) {
				assert.Equal(h.T(), 200, actualRes.StatusCode)
			},
		},
	}

	articleHandler := article_handler.NewArticleHandler()
	for _, c := range cases {
		h.Run(c.caseName, func() {
			h.routerGroup.Post("/", func(fctx *fiber.Ctx) error {
				c.expectedLogic(fctx, c)
				return articleHandler.CreateArticle(fctx)
			})

			url := "/api/v1/article"
			req := httptest.NewRequest("POST", url, strings.NewReader(c.json))
			resp, _ := h.app.Test(req, 1)
			c.expectedResponse(resp)
		})
	}
}
