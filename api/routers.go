package api

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/yasuaki640/go-intermediate-playground/api/middlewares"
	"github.com/yasuaki640/go-intermediate-playground/controllers"
	"github.com/yasuaki640/go-intermediate-playground/services"
)

func NewRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()

	r.Use(middlewares.LoggingMiddleware)

	ser := services.NewMyApService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r.Post("/article", aCon.PostArticleHandler)
	r.Get("/article/list", aCon.ArticleListHandler)
	r.Get("/article/{id:[0-9]+}", aCon.ArticleDetailHandler)
	r.Put("/article/{id:[0-9]+}/nice", aCon.PostNiceHandler)
	r.Post("/comment", cCon.PostCommentHandler)

	return r
}
