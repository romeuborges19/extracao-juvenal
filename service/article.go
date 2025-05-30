package service

import (
	"database/sql"
	"log"
	"time"
)

type Article struct {
	Date    time.Time
	Content string
	Source  string
}

func NewArticle(dateStr string, content string, source string) *Article {
	date, err := time.Parse("2006-01-02", dateStr)

	if err != nil {
		log.Fatal(err)
	}

	return &Article{Date: date, Content: content, Source: source}
}

func SaveArticle(a *Article, db *sql.DB) int {
	stmt := `INSERT INTO juvenal_news(date, content, source) VALUES ($1, $2, $3)`
	// var id int

	_, err := db.Exec(stmt, a.Date, a.Content, a.Source)
	if err != nil {
		log.Fatal(err)
	}

	return 1
}
