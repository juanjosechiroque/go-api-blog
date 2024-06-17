package tests

import (
	"go-api-blog/pkg/routes"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
)

func TestListAllPost(t *testing.T) {

	router := mux.NewRouter()

	allRoutes := routes.GetAllRoutes()

	for _, route := range allRoutes {
		router.Handle(route.Path, route.Handler).Methods(route.Method)
	}

	server := httptest.NewServer(router)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/posts").Expect().Status(200)
}

func TestGetPost(t *testing.T) {

	router := mux.NewRouter()

	allRoutes := routes.GetAllRoutes()

	for _, route := range allRoutes {
		router.Handle(route.Path, route.Handler).Methods(route.Method)
	}

	server := httptest.NewServer(router)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/posts/1").Expect().Status(200)
}

func TestCreatePost(t *testing.T) {
	router := routes.CreateAppRouter()

	server := httptest.NewServer(router)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.POST("/posts").Expect().Status(200)
}
