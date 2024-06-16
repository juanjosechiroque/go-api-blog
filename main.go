package main

import (
	"fmt"
	"go-api-blog/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	allRoutes := routes.GetAllRoutes()

	for _, route := range allRoutes {
		router.Handle(route.Path, route.Handler).Methods(route.Method)
	}

	fmt.Println("Starting")

	http.ListenAndServe(":3000", router)
}
