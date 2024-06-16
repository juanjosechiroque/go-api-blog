package routes

import (
	"go-api-blog/pkg/models"
)

func GetAllRoutes() []models.Route {
	var allRoutes []models.Route

	allRoutes = append(allRoutes, PostRoutes()...)

	return allRoutes
}
