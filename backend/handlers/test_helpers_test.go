package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func newJSONRequest(method string, target string, body string) (*gin.Context, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	request := httptest.NewRequest(method, target, bytes.NewBufferString(body))
	request.Header.Set("Content-Type", "application/json")
	context.Request = request

	return context, recorder
}

func assertStatus(recorder *httptest.ResponseRecorder, expected int) bool {
	return recorder.Code == expected || recorder.Code == http.StatusOK && expected == http.StatusOK
}
