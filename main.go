package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	api      = "/api/v1/"
	feedPath = api + "feed/"
)

var app App

func main() {

	algolia := AlgoliaConnection{}

	algolia.InitDatabseConnection()

	app = App{

		&algolia,
	}

	rssDataSlice := []RSSData{}

	rssDataChan := make(chan []RSSData)
	rssDataChan2 := make(chan []RSSData)
	rssErrorChan1 := make(chan error)
	rssErrorChan2 := make(chan error)

	log.Println("fetching RSS feed...")

	go GetRSSFeeds("http://feeds.bbci.co.uk/news/world/rss.xml", rssDataChan, rssErrorChan1)
	go GetRSSFeeds("http://rss.cnn.com/rss/edition_world.rss", rssDataChan2, rssErrorChan2)

	rssDataSlice = append(rssDataSlice, <-rssDataChan...)

	rssDataSlice = append(rssDataSlice, <-rssDataChan2...)

	log.Println(rssDataSlice)

	err := algolia.AddRecords(rssDataSlice)

	log.Println("Adding RSS records to database...")

	if err != nil {
		fmt.Println("keys", *algolia.client, goDotEnvVariable("ADMIN_API_KEY"))
		log.Fatalln("[ERROR] main ->", err.Error())
	}

	startServer()

}

func startServer() {

	http.HandleFunc(feedPath, app.handleFeedSearch)

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
