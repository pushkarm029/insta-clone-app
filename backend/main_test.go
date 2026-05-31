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
