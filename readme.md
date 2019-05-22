# Ebay Play
Basic scraper that takes a tsv file of titles and categories as input.
The scraper queries the Ebay API for each product title.
Upon recieveing a response, we concatenate the product title and category coming out of the API and compare that to the title and category from the TSV file. If these match, then we conclude that the product is correctly identified. 

### Prerequisities
Go
Ebay Production API ID

### Usage
```
go get github.com/valyala/tsvreader
cd ebay-api-retrieval
go build .
./ebay-api-retrieval
```

### CLI arguments
app_id required
```
Usage of ./ebay-play:
  -a 
      String: Application ID from Ebay
  -i
      String: Path to TSV Input File (default "data/in.tsv")
  -o
      String: Path to XML Output File (default "out/dump.xml")
  -r
      Boolean, whether to remove Header from Input (default true)
```

### To Do
Better management of fails. Retries.
Add Channels, multi thread scraping

