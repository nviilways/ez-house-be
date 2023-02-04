package testutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"git.garena.com/sea-labs-id/batch-05/adithya-kurniawan/final-project/house-booking-be/server"
	"github.com/gin-gonic/gin"
)

func ServeReq(cfg *server.RouterConfig, req *http.Request) (*gin.Engine, *httptest.ResponseRecorder) {
	router := server.NewRouter(cfg)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return router, rec
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}
