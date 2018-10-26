package main

import 
(
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func main() {

	// SiteMapIndex
	var s SiteMapIndex

	// News
	var n News

	// news_map holds string of type NewsMap{keyword, location}
	news_map := make(map[string]NewsMap)

	// Get sitemap of washington post
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")

	// Read the body
	bytes, _ := ioutil.ReadAll(resp.Body)

	// Get all sitemap urls
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)

	// Loop through each sitemap location
	for _, Location := range s.Locations {

		// Visit the sitemap
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)

		// Map xml to news
		xml.Unmarshal(bytes, &n)
		
		for idx, _ := range n.Titles{
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}

	fmt.Println("hello")
}