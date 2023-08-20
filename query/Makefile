generate-mock:
	mockgen -source="internal/service/article.go" -destination="mock/service/article.go" -package="service_mock"
	mockgen -source="internal/repository/article.go" -destination="mock/repository/article.go" -package="repository_mock"
	mockgen -source="internal/handler/article.go" -destination="mock/handler/article.go" -package="handler_mock"


run-unittest-cvg:
	echo "Run Test With Coverage"
	go test -covermode=atomic -coverprofile=coverage-reports.out ./internal/handler/... ./internal/repository/... ./internal/service/...
	echo "Calculate Coverage"
	go tool cover -func coverage-reports.out | grep total
	rm -f coverage-reports.out

.PHONY: generate-mock run-unittest-cvg