package routes

import (
	"go-api-blog/pkg/models"

	"github.com/gorilla/mux"
)

func CreateAppRouter() *mux.Router {
	router := mux.NewRouter()

	allRoutes := getAllRoutes()

	for _, route := range allRoutes {
		router.Handle(route.Path, route.Handler).Methods(route.Method)
	}

	return router
}

func getAllRoutes() []models.Route {
	var allRoutes []models.Route

	allRoutes = append(allRoutes, PostRoutes()...)

	return allRoutes
}
