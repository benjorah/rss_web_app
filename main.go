//A web app that allows user search a database for RSS feeds matching a particular topic
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
)

const (
	api      = "/api/v1/"
	feedPath = api + "feed/"
)

var app App

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	/*
		This first two if blocks are used for genearting CPU and memory profiles for the application
		We can generate this by running
		`go test -cpuprofile cpu.prof2 -memprofile mem.prof2 -bench .` - to generate profiles, and then,
		`go tool pprof --pdf ~/go/src/rss_web_app/cpu.prof2 > cpu.pdf2` - to convert profiles to pdf

	*/
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
	//Performance profile code ends here

	//main program starts
	algolia := AlgoliaConnection{}

	algolia.InitDatabseConnection()

	app = App{

		&algolia,
	}

	var errorFromChan error = nil
	rssDataSlice := []RSSData{}
	rssDataChan := make(chan []RSSData)
	rssErrorChan := make(chan error)

	RSSFeedUrls := [2]string{"http://feeds.bbci.co.uk/news/world/rss.xml", "http://rss.cnn.com/rss/edition_world.rss"}

	log.Println("fetching RSS feed...")

	//Fetch the RSS feeds from 2 sources using 2 gouroutines
	for _, url := range RSSFeedUrls {

		fmt.Println("outer running with url of " + url)
		url := url

		//we use an anonymous function to make GetRSSFeeds not depend on channels thereby making it more testable
		go func(inputPathOrString string) {

			dataSlice, err := GetRSSFeeds(inputPathOrString)

			if err != nil {
				rssErrorChan <- fmt.Errorf("while fetching from %s <= %s", url, err.Error())
				rssDataChan <- []RSSData{}
				return
			}

			rssDataChan <- dataSlice
			rssErrorChan <- nil
			return
		}(url)

		//Wait for Data and Error
		rssDataSlice = append(rssDataSlice, <-rssDataChan...)
		errorFromChan = <-rssErrorChan

	}

	if errorFromChan != nil {

		log.Println("[ERROR] main() " + errorFromChan.Error())

	}

	fmt.Println(rssDataSlice)

	// err := algolia.AddRecords(rssDataSlice)

	// log.Println("Adding RSS records to database...")

	// if err != nil {
	// 	log.Fatalln("[ERROR] main() <=", err.Error())
	// }

	//start the server
	startServer()

}

func startServer() {

	http.HandleFunc(feedPath, app.handleFeedSearch)

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}

	log.Println("Listening on port 8000")
}
