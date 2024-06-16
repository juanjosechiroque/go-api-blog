package controllers

import "net/http"

func GetAllPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/GET get all posts"))
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/GET get post"))
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/POST create post"))
}
