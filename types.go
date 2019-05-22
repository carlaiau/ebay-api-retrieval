package main

import "encoding/xml"

type Document struct {
	ID         string
	Price      float32
	Title      string
	Category   string
	PictureURL string
}

type ScrapedListing struct {
	ItemID             string `xml:"id"`
	EbayID             string `xml:"ebayID"`
	Title              string `xml:"title"`
	Description        string `xml:"description"`
}

type SingleListing struct {
	Description string `xml:"Item>Description"`
	PictureURL  string `xml:"Item>PictureURL"`
}
type Item struct {
	ItemID              string  `xml:"itemId"`
	Title               string  `xml:"title"`
	CurrentPrice        float64 `xml:"sellingStatus>currentPrice"`
	PrimaryCategoryName string  `xml:"primaryCategory>categoryName"`
}

type ResponseXML struct {
	XMLName   xml.Name `xml:"findItemsByKeywordsResponse"`
	Items     []Item   `xml:"searchResult>item"`
	TimeStamp string   `xml:"timestamp"`
}

type ExpiredResponseXML struct {
	XMLName   xml.Name `xml:"findCompletedItemsResponse"`
	Items     []Item   `xml:"searchResult>item"`
	TimeStamp string   `xml:"timestamp"`
}

