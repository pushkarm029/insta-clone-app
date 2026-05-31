package handlers

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUploadPostToFirestoreRejectsInvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	context, recorder := newJSONRequest(http.MethodPost, "/api/upload/test@example.com", "{")

	err := UploadPostToFirestore(context.Request.Context(), nil, context, "test@example.com")

	if err == nil {
		t.Fatal("expected invalid JSON to return an error")
	}
	if !assertStatus(recorder, http.StatusBadRequest) {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}
