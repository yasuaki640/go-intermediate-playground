package repositories

import (
	"database/sql"
	"fmt"
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
		newArticle.CreatedAt = createdTime.Time
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

const PerPage = 5

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?;
	`
	rows, err := db.Query(sqlStr, PerPage, (page-1)*PerPage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}

	return articleArray, nil
}

func UpdateNiceNum(db *sql.DB, articleId int) error {
	const sqlStr = `
		update articles set nice = nice + 1 where article_id = ?;
	`
	_, err := db.Exec(sqlStr, articleId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func SelectCommentList(db *sql.DB, articleId int) ([]models.Comment, error) {
	const sqlStr = `
		select comment_id, article_id, message, created_at
		from comments
		where article_id = ?;
	`
	rows, err := db.Query(sqlStr, articleId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &comment.CreatedAt)
		if err != nil {
			fmt.Println(err)
		} else {
			commentArray = append(commentArray, comment)
		}

	}
	return commentArray, err
}

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlINsertStr = `
		insert into comments (article_id, message, created_at)
		values (?, ?, now());
	`
	result, err := db.Exec(sqlINsertStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	const sqlGetStr = `
		select comment_id, article_id, message, created_at from comments
		where comment_id = ?;
	`
	id, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	row := db.QueryRow(sqlGetStr, id)
	if err := row.Err(); err != nil {
		return models.Comment{}, err
	}

	var newComment models.Comment
	var createdTime sql.NullTime
	err = row.Scan(&newComment.CommentID, &newComment.ArticleID, &newComment.Message, &createdTime)
	if err != nil {
		return models.Comment{}, err
	}
	if createdTime.Valid {
		newComment.CreatedAt = createdTime.Time
	}

	return newComment, err
}
