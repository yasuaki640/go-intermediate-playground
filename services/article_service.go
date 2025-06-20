package services

import (
	"database/sql"
	"errors"
	"github.com/yasuaki640/go-intermediate-playground/apperrors"
	"github.com/yasuaki640/go-intermediate-playground/models"
	"github.com/yasuaki640/go-intermediate-playground/repositories"
)

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to insert article")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "failed to get article")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "failed to get article")
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if errors.Is(err, sql.ErrNoRows) {
		err = apperrors.NAData.Wrap(err, "no data")
		return nil, err
	}
	err = apperrors.GetDataFailed.Wrap(err, "failed to get article list")
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}

	return articleList, nil
}
