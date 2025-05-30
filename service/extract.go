package service

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func formatDateString(date string) string {
	return date[:4] + "-" + date[4:6] + "-" + date[6:8]
}

func readNewsContent(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	s := bufio.NewScanner(f)

	var content string

	if s.Scan() {
		for s.Scan() {
			line := s.Text()
			if line != "" {
				content = content + " " + line
			}
		}
	}

	return content
}

func treatNewsFile(f os.DirEntry, dir string, source string) *Article {
	newsFile := f.Name()

	parts := strings.Split(newsFile, ".")
	if parts[1] != "txt" {
		errors.New(".txt file expected")
	}

	dateStr := parts[0]
	dateStr = formatDateString(dateStr[:8])

	content := readNewsContent(dir + newsFile)
	return NewArticle(dateStr, content, source)
}

func extractFolder(source string, files []os.DirEntry, root string, db *sql.DB, wg *sync.WaitGroup) {
	for _, f := range files {
		if f.IsDir() == false {
			continue
		}

		fileName := f.Name()
		subDir := root + fileName + "/"

		news, err := os.ReadDir(subDir)
		if err != nil {
			panic(err)
		}

		for _, n := range news {
			a := treatNewsFile(n, subDir, source)
			SaveArticle(a, db)
		}
	}
	wg.Done()
	fmt.Printf("Dados de %v extraídos com sucesso\n", source)
}

func CollectNews(files []os.DirEntry, root string, db *sql.DB) {
	var wg sync.WaitGroup

	for _, file := range files {
		source := file.Name()
		subDir := root + source + "/"

		months, err := os.ReadDir(subDir)
		if err != nil {
			panic(err)
		}

		wg.Add(1)
		go extractFolder(source, months, subDir, db, &wg)
	}
	wg.Wait()
}
