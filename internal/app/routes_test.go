package app_test

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app"
	"github.com/stretchr/testify/suite"
)

type RouteTestSuites struct {
	suite.Suite
	client *http.Client
	srv    *httptest.Server
}

func (rts *RouteTestSuites) readResponse(res *http.Response) string {
	resBodyBytes, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	return string(resBodyBytes)
}

func (rts *RouteTestSuites) SetupTest() {
	rts.client = &http.Client{}

	err := godotenv.Load(".env.testing")
	rts.NoError(err)

	_app, err := app.InitializeFromEnv()
	rts.NoError(err)

	ec := echo.New()
	ec = _app.Routes(ec)
	rts.srv = httptest.NewServer(ec.Server.Handler)
}

func (rts *RouteTestSuites) TearDownTest() {
	rts.srv.Close()
}

func (rts *RouteTestSuites) TestRoutes_AbleToHitHealthCheck() {

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/v1/health", rts.srv.URL),
		nil,
	)
	res, err := rts.client.Do(req)
	resBody := rts.readResponse(res)
	rts.NoError(err)

	jsonBody := map[string]interface{}{}
	err = json.Unmarshal([]byte(resBody), &jsonBody)
	rts.NoError(err)

	rts.Equal(http.StatusOK, res.StatusCode)
	rts.Equal("ok", jsonBody["status"])
}

func TestRouteTestSuites(t *testing.T) {
	suite.Run(t, new(RouteTestSuites))
}
