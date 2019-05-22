package main

import (
	"fmt"
	"github.com/valyala/tsvreader"
	"os"
	"strings"
)

func getCategoryFromTree(t string) string {
	categories := strings.Split(t, "> ")
	return categories[len(categories)-1]

}
func getDocs(filePath string, removeHeader bool) []Document {
	parsedHeader := false
	var documents []Document

	data, openErr := os.Open(filePath)
	if openErr != nil {
		fmt.Println("Open Error")
		panic(openErr)
	}

	r := tsvreader.New(data)

	for r.Next() {
		if !parsedHeader && removeHeader{
			for i := 0; i < 5; i++ {
				r.SkipCol()
			}
			parsedHeader = true
		} else {

			doc := Document{
				ID:         r.String(),
				Title:      strings.Trim(r.String(), " "),
				Category:   getCategoryFromTree(r.String()),
			}
			documents = append(documents, doc)
		}

	}

	if parseErr := r.Error(); parseErr != nil {
		fmt.Println("Parse Error")
		panic(parseErr)

	}

	return documents

}
