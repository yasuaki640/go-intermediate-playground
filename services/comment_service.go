package services

import (
	"database/sql"
	"errors"
	"github.com/yasuaki640/go-intermediate-playground/apperrors"
	"github.com/yasuaki640/go-intermediate-playground/models"
	"github.com/yasuaki640/go-intermediate-playground/repositories"
)

func (s *MyAppService) PostNiceService(articleID int) (models.Article, error) {

	err := repositories.UpdateNiceNum(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "failed to update nice num")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "failed to update nice num")
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
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert comment")
		return models.Comment{}, err
	}

	return newComment, nil
}
