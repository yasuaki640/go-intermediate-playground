package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/yasuaki640/go-intermediate-playground/models"
	"github.com/yasuaki640/go-intermediate-playground/services"
	"net/http"
	"strconv"
)

type MyAppController struct {
	services *services.MyAppService
}

func NewMyAppController(services *services.MyAppService) *MyAppController {
	return &MyAppController{services: services}
}

func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := c.services.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	pageStr := req.URL.Query().Get("page")

	var page int

	if pageStr == "" {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	}

	artcles, err := c.services.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(artcles)
}

// func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	article, err := c.services.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "failed to get article", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	article, err := c.services.PostNiceService(articleID)
	if err != nil {
		http.Error(w, "failed to exec", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment, err := c.services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
