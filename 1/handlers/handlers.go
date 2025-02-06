package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/yasuaki640/go-intermediate-playground/models"
	"io"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from Hell!!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	articles := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	articles := []models.Article{models.Article1, models.Article2}
	var resData *models.Article
	for _, article := range articles {
		if article.ID == articleID {
			resData = &article
		}
	}

	jsonData, err := json.Marshal(&resData)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
