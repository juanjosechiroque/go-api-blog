package routes

import (
	"go-api-blog/pkg/models"
)

func GetAllRoutes() []models.Route {
	var allRoutes []models.Route

	postRoutes := PostRoutes()
	allRoutes = append(allRoutes, postRoutes...)

	return allRoutes
}
