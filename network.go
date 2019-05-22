package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func createSearch(keywords string) string {

	endpoint, _ := url.Parse("http://svcs.ebay.com/services/search/FindingService/v1")
	params := url.Values{}
	params.Add("OPERATION-NAME", "findItemsByKeywords")
	params.Add("SERVICE-VERSION", "1.0.0")
	params.Add("SECURITY-APPNAME", appID)
	params.Add("keywords", keywords)
	params.Add("GLOBAL-ID", "EBAY-US")
	params.Add("RESPONSE-DATA-FORMAT", "XML")
	params.Add("REST-PAYLOAD", "")

	endpoint.RawQuery = params.Encode()

	return endpoint.String()
}

func createExpiredSearch(keywords string) string{
	endpoint, _ := url.Parse("http://svcs.ebay.com/services/search/FindingService/v1?")
	params := url.Values{}
	params.Add("OPERATION-NAME", "findCompletedItems")
	params.Add("SERVICE-VERSION", "1.7.0")
	params.Add("SECURITY-APPNAME", appID)
	params.Add("keywords", keywords)
	params.Add("RESPONSE-DATA-FORMAT", "XML")
	params.Add("REST-PAYLOAD", "")
	endpoint.RawQuery = params.Encode()
	return endpoint.String()
}

func createSingleLookup(id string) string {
	endpoint, _ := url.Parse("http://open.api.ebay.com/shopping")
	params := url.Values{}
	params.Add("callname", "GetSingleItem")
	params.Add("responseencoding", "XML")
	params.Add("appid", appID)
	params.Add("siteid", "0")
	params.Add("version", "967")
	params.Add("ItemID", id)
	params.Add("IncludeSelector", "Description")

	endpoint.RawQuery = params.Encode()

	return endpoint.String()
}

func getSearchResponse(url string) ResponseXML {
	r, _ := http.Get(url)
	// // Need to handle Errors here to account for loss of comms
	response, _ := ioutil.ReadAll(r.Body)
	v := ResponseXML{}
	err := xml.Unmarshal([]byte(response), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return v
}


func getExpiredSearchResponse(url string) ExpiredResponseXML {
	r, _ := http.Get(url)
	// // Need to handle Errors here to account for loss of comms
	response, _ := ioutil.ReadAll(r.Body)
	v := ExpiredResponseXML{}
	err := xml.Unmarshal([]byte(response), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return v
}


func getSingleResponse(url string) SingleListing {
	r, _ := http.Get(url)
	// // Need to handle Errors here to account for loss of comms
	response, _ := ioutil.ReadAll(r.Body)
	v := SingleListing{}
	err := xml.Unmarshal([]byte(response), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return v
}
