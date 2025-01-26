package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/yasuaki640/go-intermediate-playground/handlers"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.HandleFunc("/", handlers.HelloHandler)
	r.Get("/article", handlers.PostArticleHandler)
	r.Get("/article/list", handlers.ArticleListHandler)
	r.Get("/article/{id}", handlers.ArticleDetailHandler)
	r.Post("/article/nice", handlers.PostNiceHandler)
	r.Post("/comment", handlers.PostCommentHandler)

	log.Println("listening at port 8080")

	err := http.ListenAndServe(":8080", r)
	log.Fatal(err) // exit 1のときにログ出力される
}
