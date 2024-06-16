package routes

import (
	"go-api-blog/pkg/controllers"
	"go-api-blog/pkg/models"
)

func PostRoutes() []models.Route {
	return []models.Route{
		{Path: "/posts", Method: "GET", Handler: controllers.GetAllPostHandler},
		{Path: "/posts/{id}", Method: "GET", Handler: controllers.GetPostHandler},
		{Path: "/posts", Method: "POST", Handler: controllers.CreatePostHandler},
	}
}
