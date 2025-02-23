package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yasuaki640/go-intermediate-playground/handlers"
	"log"
	"net/http"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.HandleFunc("/", handlers.HelloHandler)
	r.Post("/article", handlers.PostArticleHandler)
	r.Get("/article/list", handlers.ArticleListHandler)
	r.Get("/article/{id}", handlers.ArticleDetailHandler)
	r.Post("/article/nice", handlers.PostNiceHandler)
	r.Post("/comment", handlers.PostCommentHandler)

	log.Println("listening at port 8080")

	err = http.ListenAndServe(":8080", r)
	log.Fatal(err) // exit 1のときにログ出力される
}
