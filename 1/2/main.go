package main

import (
	"github.com/yasuaki640/go-intermediate-playground/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("listening at port 8080")

	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err) // exit 1のときにログ出力される
}
