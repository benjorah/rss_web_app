package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/mmcdole/gofeed"
)

//ReadAndParse takes an input string, which could be a url, raw rss string or file path
//and uses it to fetch and parse rss feeds
func ReadAndParse(inputPathOrString string) (rssFeed []*gofeed.Item, funcErr error) {

	var feed *gofeed.Feed
	var parseError error

	feedParser := gofeed.NewParser()

	if strings.HasPrefix(strings.ToLower(inputPathOrString), "http") {
		feed, parseError = feedParser.ParseURL(inputPathOrString)

	} else if strings.HasSuffix(strings.ToLower(inputPathOrString), ".xml") {

		file, err := os.Open(inputPathOrString)
		defer file.Close()

		if err != nil {

			return nil, fmt.Errorf("ReadAndParse : an error was encountered while attemptimg to open this file ; %#v" + err.Error())
		}
		feed, parseError = feedParser.Parse(file)

	} else {

		feed, parseError = feedParser.ParseString(inputPathOrString)

	}

	if parseError != nil {

		return nil, fmt.Errorf("ReadAndParse : an error was encountered while attemptimg to parse the RSS feeds ; %#v" + parseError.Error())

	}

	return feed.Items, funcErr

}

//GetRSSFeeds takes an input string, which could be a url, raw rss string or file path
//and uses it to fetch and parse rss feeds
func GetRSSFeeds(inputPathOrString string, dataChan chan<- []RSSData, errorChan chan<- error) {

	innerWaitGroup := &sync.WaitGroup{}
	rssDataSlice := []RSSData{}

	feed, _ := ReadAndParse(inputPathOrString)

	transformedDataChan := make(chan RSSData)

	for _, element := range feed {

		innerWaitGroup.Add(1)
		go transFormRSSFeedToRSSData(element, innerWaitGroup, transformedDataChan)
		rssDataSlice = append(rssDataSlice, <-transformedDataChan)

	}

	innerWaitGroup.Wait()

	// fmt.Println(rssDataSlice)

	dataChan <- rssDataSlice
	close(dataChan)

}

func transFormRSSFeedToRSSData(feed *gofeed.Item, wg *sync.WaitGroup, dataChan chan<- RSSData) {

	transformedRSS := RSSData{
		feed.Title,
		feed.Description,
		feed.Link,
		feed.PublishedParsed,
	}

	dataChan <- transformedRSS

	wg.Done()

}
