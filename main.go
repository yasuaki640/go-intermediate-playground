package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yasuaki640/go-intermediate-playground/controllers"
	"github.com/yasuaki640/go-intermediate-playground/routers"
	"github.com/yasuaki640/go-intermediate-playground/services"
	"log"
	"net/http"
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

	ser := services.NewMyApService(db)
	con := controllers.NewMyAppController(ser)

	r := routers.NewRouter(con)

	log.Println("listening at port 8080")

	err = http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
