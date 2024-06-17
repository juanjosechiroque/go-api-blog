package main

import (
	"fmt"
	"go-api-blog/pkg/routes"
	"net/http"
)

func main() {
	router := routes.CreateAppRouter()

	fmt.Println("Running on port 3001")

	http.ListenAndServe(":3001", router)
}
