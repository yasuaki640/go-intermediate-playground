package services

import "database/sql"

type MyAppService struct {
	db *sql.DB
}

func NewMyApService(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
