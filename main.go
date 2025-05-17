package main

import (
	"fmt"
	s "main/service"
	"os"
	"time"
)

const datasetsPath = "/home/rborges/projetos/tcc/datasets/"

func main() {
	const newsRoot = datasetsPath + "news_db/"

	files, err := os.ReadDir(newsRoot)
	if err != nil {
		panic(err)
	}

	db, err := OpenConn()

	start := time.Now()
	s.CollectNews(files, newsRoot, db)
	fmt.Printf("Tempo: %v", time.Since(start))
}
