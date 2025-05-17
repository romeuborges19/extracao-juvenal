package main

import (
	"fmt"
	"log"
	s "main/service"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func getRoot() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("DATASET_ROOT")
}

func main() {
	newsRoot := getRoot() + "news_db/"

	files, err := os.ReadDir(newsRoot)
	if err != nil {
		panic(err)
	}

	db, err := OpenConn()

	start := time.Now()
	s.CollectNews(files, newsRoot, db)
	fmt.Printf("Tempo: %v", time.Since(start))
}
