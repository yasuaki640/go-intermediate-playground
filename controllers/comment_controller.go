package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/yasuaki640/go-intermediate-playground/apperrors"
	"github.com/yasuaki640/go-intermediate-playground/controllers/services"
	"github.com/yasuaki640/go-intermediate-playground/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "failed to decode json")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
