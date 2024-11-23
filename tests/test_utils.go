package tests

import (
	"bytes"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func SetupTestRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func PerformRequest(router *gin.Engine, method, path string, body []byte) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
