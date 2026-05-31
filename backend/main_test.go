package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInitRoutesRegistersAPIEndpoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	initRoutes(router)

	routes := map[string]bool{}
	for _, route := range router.Routes() {
		routes[route.Method+" "+route.Path] = true
	}

	expectedRoutes := []string{
		"GET /api/explore/posts",
		"GET /api/search/users",
		"POST /api/upload/:userID",
		"POST /api/like/:OverAcEmail/:OverAcImages",
		"POST /api/comment/post/:OverAcEmail/:OverAcImages",
		"POST /api/follow/:OverAcEmail/:UserID",
		"GET /api/profile/user/:userID",
		"GET /api/comment/get/:OverAcEmail/:urlModEncoder",
		"GET /api/home/:userID",
	}

	for _, expectedRoute := range expectedRoutes {
		if !routes[expectedRoute] {
			t.Fatalf("expected route %q to be registered", expectedRoute)
		}
	}
}

func TestAllowedOriginsUsesConfiguredList(t *testing.T) {
	t.Setenv("CORS_ALLOWED_ORIGINS", "http://localhost:3000, https://example.com")

	origins := allowedOrigins()

	if len(origins) != 2 {
		t.Fatalf("expected 2 origins, got %d", len(origins))
	}
	if origins[0] != "http://localhost:3000" {
		t.Fatalf("expected first origin to be trimmed, got %q", origins[0])
	}
	if origins[1] != "https://example.com" {
		t.Fatalf("expected second origin to be trimmed, got %q", origins[1])
	}
}

func TestServerAddressDefaultsTo8080(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("BACKEND_PORT", "")

	if address := serverAddress(); address != ":8080" {
		t.Fatalf("expected default address :8080, got %q", address)
	}
}
