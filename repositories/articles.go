package repositories

import (
	"database/sql"
	"github.com/yasuaki640/go-intermediate-playground/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlInsertStr = `
		insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());
	`
	result, err := db.Exec(sqlInsertStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	const sqlGetStr = `
		select * from articles where article_id = ?;
	`
	id, err := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}

	row := db.QueryRow(sqlGetStr, id)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var newArticle models.Article
	var createdTime sql.NullTime
	err = row.Scan(&newArticle.ID, &newArticle.Title, &newArticle.Contents, &newArticle.UserName, &newArticle.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return newArticle, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}
