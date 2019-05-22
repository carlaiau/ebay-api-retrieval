package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"flag"
)

var appID string = ""
var listingsFound int = 0
var outputFilePath string = ""

func getSingleItem(doc Document, item Item, single SingleListing) bool {
	if (doc.Title == strings.Trim(item.Title, " ")) && (doc.Category == strings.Trim(item.PrimaryCategoryName, " ")) {
		listing := ScrapedListing{
			ItemID:             doc.ID,
			EbayID:             item.ItemID,
			Title:              strings.Trim(doc.Title, " "),
			Description:        single.Description,
		}
		outputFile, err := os.OpenFile(outputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		output, mErr := xml.MarshalIndent(listing, "  ", "    ")
		if mErr != nil {
			panic(mErr)
		}
		outputFile.Write(output)
		return true
	}
	return false
}

func expiredSearch(doc Document){
	searchUrl := createExpiredSearch(doc.Title)
	v := getExpiredSearchResponse(searchUrl)
	if len(v.Items) > 0 {
		item := v.Items[0]
		singleUrl := createSingleLookup(item.ItemID)
		single := getSingleResponse(singleUrl)

		foundCorrectData := getSingleItem(doc, item, single)

		if foundCorrectData {
			listingsFound++
		}
	}

}
func getData(docs []Document) {

	for _, doc := range docs {
		searchUrl := createSearch(doc.Title)
		v := getSearchResponse(searchUrl)
		if len(v.Items) > 0 {
			item := v.Items[0]

			singleUrl := createSingleLookup(item.ItemID)
			single := getSingleResponse(singleUrl)

			foundCorrectData := getSingleItem(doc, item, single)

			if foundCorrectData {
				listingsFound++
			} else{
				expiredSearch(doc)
			}
		} else { 
			expiredSearch(doc)
			
		}

	}

	fmt.Printf("Scrape Completed\n%d from %d listings found", listingsFound, len(docs))
}

func main() {
	
	flag.StringVar(&appID, "a", "", "String: Application ID from Ebay")
	flag.StringVar(&outputFilePath, "o", "out/dump.xml", "Path to XML Output File")
	inputFilePath := flag.String("i", "data/in.tsv", "Path to TSV Input File")
	removeHeader := flag.Bool("r", false, "Boolean, whether to remove Header from Input")
	
	flag.Parse()

	docs := getDocs(*inputFilePath, *removeHeader)	
	getData(docs)


	
}
