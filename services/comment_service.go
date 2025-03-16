package services

import (
	"github.com/yasuaki640/go-intermediate-playground/models"
	"github.com/yasuaki640/go-intermediate-playground/repositories"
)

func (s *MyAppService) PostNiceService(articleID int) (models.Article, error) {

	err := repositories.UpdateNiceNum(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
