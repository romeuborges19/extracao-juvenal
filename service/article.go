package service

import (
	"database/sql"
	"log"
	"time"
)

type Article struct {
	Date    time.Time
	Content string
}

func NewArticle(dateStr string, content string) *Article {
	date, err := time.Parse("2006-01-02", dateStr)

	if err != nil {
		log.Fatal(err)
	}

	return &Article{Date: date, Content: content}
}

func SaveArticle(a *Article, db *sql.DB) int {
	stmt := `INSERT INTO juvenal_news(date, content) VALUES ($1, $2)`
	// var id int

	_, err := db.Exec(stmt, a.Date, a.Content)
	if err != nil {
		log.Fatal(err)
	}

	return 1
}
