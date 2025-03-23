package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/yasuaki640/go-intermediate-playground/controllers"
)

func NewRouter(con *controllers.MyAppController) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/article", con.PostArticleHandler)
	r.Get("/article/list", con.ArticleListHandler)
	r.Get("/article/{id:[0-9]+}", con.ArticleDetailHandler)
	r.Put("/article/{id:[0-9]+}/nice", con.PostNiceHandler)
	r.Post("/comment", con.PostCommentHandler)

	return r
}
