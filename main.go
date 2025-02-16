package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yasuaki640/go-intermediate-playground/models"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	const sqlStr = `
		select *
		from articles
		where article_id = ?
	`
	row := db.QueryRow(sqlStr, 10000)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var article models.Article
	var createdTime sql.NullTime
	err = row.Scan(
		&article.ID,
		&article.Title,
		&article.Contents,
		&article.UserName,
		&article.NiceNum,
		&createdTime,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Println("%+v\n", article)
}
