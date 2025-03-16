package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	controllers "github.com/yasuaki640/go-intermediate-playground/handlers"
	"github.com/yasuaki640/go-intermediate-playground/services"
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

	ser := services.NewMyApService(db)
	con := controllers.NewMyAppController(ser)

	r := chi.NewRouter()

	r.Post("/article", con.PostArticleHandler)
	r.Get("/article/list", con.ArticleListHandler)
	r.Get("/article/{id}", con.ArticleDetailHandler)
	r.Put("/article/{id}/nice", con.PostNiceHandler)
	r.Post("/comment", con.PostCommentHandler)

	log.Println("listening at port 8080")

	err = http.ListenAndServe(":8080", r)
	log.Fatal(err) // exit 1のときにログ出力される
}
