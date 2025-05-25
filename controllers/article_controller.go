package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/yasuaki640/go-intermediate-playground/apperrors"
	"github.com/yasuaki640/go-intermediate-playground/controllers/services"
	"github.com/yasuaki640/go-intermediate-playground/models"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "failed to decode json")
		apperrors.ErrorHandler(w, req, err)
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	pageStr := req.URL.Query().Get("page")

	var page int
	if pageStr == "" {
		page = 1
	} else {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "failed to convert page")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	}

	artcles, err := c.service.GetArticleListService(page)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(artcles)
}

func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "failed to convert articleID")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "failed to convert articleID")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.PostNiceService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(article)
}
