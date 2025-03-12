package services

import (
	"github.com/yasuaki640/go-intermediate-playground/models"
	"github.com/yasuaki640/go-intermediate-playground/repositories"
)

func PostArticleService(article models.Article) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	var newArticle models.Article
	newArticle, err = repositories.InsertArticle(db, article)
	if err != nil {
		return models.Article{}, err
	}

	return newArticle, nil
}

func GetArticleService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func GetArticleListService(page int) ([]models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	articleList, err := repositories.SelectArticleList(db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func PostNiceService(articleID int) (models.Article, error) {
	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	err = repositories.UpdateNiceNum(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	return article, nil
}
