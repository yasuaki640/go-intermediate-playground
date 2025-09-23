package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yasuaki540/go-intermediate-playground/handlers"
)

func main() {

	r := chi.NewRouter()

	r.Post("/article", handlers.PostArticleHandler)
	r.Get("/article/list", handlers.ArticleListHandler)
	r.Get("/article/1", handlers.ArticleDetailHandler)
	r.Post("/article/nice", handlers.PostNiceHandler)
	r.Post("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
