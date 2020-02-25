package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/mmcdole/gofeed"
)

//ReadAndParse fetches and parses RSS feeds..returning a slice of Items ( a custome type defined in the package `github.com/mmcdole/gofeed`).
//It takes a string representing one of the following;
//1.) A path to an xml file containing RSS Feed
//2.) A Url to fetch RSS Feed
//3.) RSS data in a string
//It returns any error it encounters.
func ReadAndParse(inputPathOrString string) (rssFeed []*gofeed.Item, err error) {

	if strings.TrimSpace(inputPathOrString) == "" {
		return nil, fmt.Errorf("[ERROR] ReadAndParse() : the input string should not be empty ")

	}

	var feed *gofeed.Feed
	var parseError error

	RSSParser := gofeed.NewParser()

	//check what class of input (RSS string, XML file path or Url)
	//Another option is to accet io.Writer interface as input but reading from the string buffer might waste unnecessary time
	//We are more likely to get strings rather than file paths
	if strings.HasPrefix(strings.ToLower(inputPathOrString), "http") {
		feed, parseError = RSSParser.ParseURL(inputPathOrString)

	} else if strings.HasSuffix(strings.ToLower(inputPathOrString), ".xml") {

		file, fileError := os.Open(inputPathOrString)
		defer file.Close()

		if fileError != nil {

			return nil, fmt.Errorf("[ERROR] ReadAndParse() with input string %s : an error was encountered while attemptimg to open this file ; %s", inputPathOrString, err.Error())
		}
		feed, parseError = RSSParser.Parse(file)

	} else {

		feed, parseError = RSSParser.ParseString(inputPathOrString)

	}

	if parseError != nil {

		return nil, fmt.Errorf("[ERROR] ReadAndParse() with input string %s : an error was encountered while attemptimg to parse the RSS feeds ; %s", inputPathOrString, parseError.Error())

	}

	if len(feed.Items) == 0 {

		return nil, fmt.Errorf("[ERROR] ReadAndParse() with input string %s : the input string is not of RSS format", inputPathOrString)

	}

	return feed.Items, err

}

//GetRSSFeeds takes an input string, which could be a url, raw rss string or file path, just like ReadAndParse above,
//and uses it to fetch and parse rss feeds.
//This function works with ReadAndParse() to parse the feeds
//and then transFormRSSFeedToCustomData() to  transfrom the feed to the appropriate type.
//It communicates through channels and is appropriate for running in a gouroutine
func GetRSSFeeds(inputPathOrString string) (data []RSSData, err error) {

	if strings.TrimSpace(inputPathOrString) == "" {
		return nil, fmt.Errorf("[ERROR] GetRSSFeeds() : the input string should not be empty ")

	}

	innerWaitGroup := &sync.WaitGroup{}
	RSSDataSlice := []RSSData{}

	feed, err := ReadAndParse(inputPathOrString)

	//if error occurs, send the error through the channel and close both error and data channels
	if err != nil {

		return nil, fmt.Errorf("[ERROR] GetRSSFeeds() with input string %s <= %s", inputPathOrString, err.Error())

	}

	//channel for communicating with transFormRSSFeedToCustomData() and it's error
	transformedDataChan := make(chan RSSData)
	transformedErrorChan := make(chan error)

	//loop through each item in the feed and start a goroutine to transform each on
	//using a goroutine allows multipls items to be transformed at a time rather than serially
	for _, element := range feed {

		innerWaitGroup.Add(1)

		//we use an anonymous function to make TransFormRSSItemToCustomData not depend on channels thereby making it more testable
		go func(feedItem *gofeed.Item, wg *sync.WaitGroup, dataChan chan<- RSSData) {

			transformedData, err := TransFormRSSItemToCustomData(feedItem)

			transformedDataChan <- transformedData
			transformedErrorChan <- err

			//complete this goroutine
			wg.Done()

		}(element, innerWaitGroup, transformedDataChan)

		RSSDataSlice = append(RSSDataSlice, <-transformedDataChan)
		err = <-transformedErrorChan

		//This error shouldn't stop the function as more feed items can be transformed
		if err != nil {

			err = fmt.Errorf("[ERROR] GetRSSFeeds() with input string %s <= %s", inputPathOrString, err)

		}

	}

	//wait for all goroutines to complete before sending the data through the channel and exiting
	innerWaitGroup.Wait()

	return RSSDataSlice, err

}

//TransFormRSSItemToCustomData transforms an RSS items to a custom data type and returns it
func TransFormRSSItemToCustomData(feedItem *gofeed.Item) (data RSSData, err error) {

	if feedItem == nil {

		return RSSData{}, fmt.Errorf("[ERROR] TransFormRSSItemToCustomData() : FeedItem should not be nil")
	}

	transformedRSS := RSSData{
		feedItem.Title,
		feedItem.Description,
		feedItem.Link,
		feedItem.PublishedParsed,
	}

	return transformedRSS, nil

}
