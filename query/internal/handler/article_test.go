package handler_test

import (
	"bou.ke/monkey"
	"bytes"
	"fmt"
	article_handler "gcnt/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"time"
)

func (h *HandlerTestSuite) TestGetArticleByID() {
	type testCase struct {
		caseName         string
		expectedLogic    func(fctx *fiber.Ctx, c testCase)
		expectedResponse func(actualRes *http.Response)
		id               string
	}
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, 1, 1, 20, 34, 58, 651387237, time.UTC)
	})

	cases := []testCase{
		{
			caseName: "Success",
			id:       "1232",
			expectedLogic: func(fctx *fiber.Ctx, c testCase) {
				ctx := fctx.Context()

				h.mockService.EXPECT().
					GetByID(ctx, c.id).
					Return(gomock.Any(), nil)
			},
			expectedResponse: func(actualRes *http.Response) {
				assert.Equal(h.T(), 200, actualRes.StatusCode)
			},
		},
	}

	articleHandler := article_handler.NewArticleHandler()
	for _, c := range cases {
		h.Run(c.caseName, func() {
			h.routerGroup.Get("/:articleID", func(fctx *fiber.Ctx) error {
				c.expectedLogic(fctx, c)
				return articleHandler.GetArticleByID(fctx)
			})

			url := fmt.Sprintf("/api/v1/article/%s", c.id)
			req := httptest.NewRequest("GET", url, bytes.NewReader(nil))
			resp, _ := h.app.Test(req, 1)
			c.expectedResponse(resp)
		})
	}
}
