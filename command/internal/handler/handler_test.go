package handler_test

import (
	"encoding/json"
	"fmt"
	"gcnt/config"
	"gcnt/internal/repository"
	"gcnt/internal/service"
	handlermock "gcnt/mock/handler"
	repository_mock "gcnt/mock/repository"
	service_mock "gcnt/mock/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type HandlerTestSuite struct {
	suite.Suite
	mockHandler    *handlermock.MockArticleHandlerImpl
	mockRepository *repository_mock.MockIArticleRepository
	mockService    *service_mock.MockIArticleService
	routerGroup    fiber.Router
	app            *fiber.App
}

func (h *HandlerTestSuite) SetupSuite() {
	mockCtrl := gomock.NewController(h.T())
	defer mockCtrl.Finish()

	//config
	config.Instance = &config.AppConfig{
		AppName:            "Command Test",
		AppEnv:             "test",
		AppMode:            "test",
		MySqlHost:          "localhost",
		MySqlPort:          3306,
		MySqlDb:            "article_db",
		MySqlUsername:      "a",
		MySqlPassword:      "a",
		RedisHost:          "localhost",
		RedisPort:          3306,
		RedisPassword:      "a",
		NatsAddress:        "localhost:4222",
		CommandServiceHost: "http://localhost",
		CommandServicePort: 0,
		QueryServiceHost:   "http://localhost",
		QueryServicePort:   0,
	}

	//repo
	h.mockRepository = repository_mock.NewMockIArticleRepository(mockCtrl)
	repository.ArticleRepositoryInstance = h.mockRepository

	//service
	h.mockService = service_mock.NewMockIArticleService(mockCtrl)
	service.ArticleServiceInstance = h.mockService

	h.app = fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "apitest",
		AppName:       "apptest",
	})

	api := h.app.Group("/api/v1")
	h.routerGroup = api.Group("/article")
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (h *HandlerTestSuite) parseResponseJson(res *http.Response) map[string]interface{} {
	var response map[string]interface{}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return response
}

func (h *HandlerTestSuite) printResult(caseName interface{}, expected interface{}, result interface{}) {
	const colorReset = "\033[0m"
	const colorBlue = "\033[34m"
	const colorGreen = "\033[32m"

	fmt.Printf(colorGreen+"%v Expected "+colorReset+" %+v \n"+
		colorBlue+"COMPARE TO"+colorReset+
		" Actual %+v \n\n\n", caseName, expected, result)
}
