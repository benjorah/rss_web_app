//A web app that allows user search a database for RSS feeds matching a particular topic
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"rss_web_app/database"
	"rss_web_app/rssreader"
	"runtime"
	"runtime/pprof"

	_ "github.com/go-sql-driver/mysql"
)

const (
	api      = "/api/v1/"
	feedPath = api + "feed"
)

var app App

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var fetchrss = flag.Bool("fetchrss", false, "fetch rss feed and store in a database")

func main() {

	flag.Parse()

	//main program starts
	// algolia := AlgoliaConnection{}
	// algolia.InitDatabseConnection()
	// app = App{
	// 	&algolia,
	// }

	mysql := database.MsqlConnection{}

	err := mysql.InitDatabseConnection()
	app = App{
		&mysql,
	}

	defer mysql.GetConnectionObject().Close()

	if err != nil {

		log.Println("[ERROR] main() <= " + err.Error())

	}

	if *cpuprofile != "" || *memprofile != "" {

		createPerformanceProfile(cpuprofile, memprofile)

	}

	if *fetchrss {
		RSSFeedUrls := []string{"http://feeds.bbci.co.uk/news/world/rss.xml", "http://rss.cnn.com/rss/edition_world.rss"}

		fetchRSSFeed(RSSFeedUrls)
	}

	//start the server
	startServer()

}

//startServer initiates action to start the http server
func startServer() {

	http.HandleFunc(feedPath, app.handleFeedSearch)

	log.Println("Starting server on port 8000...")

	err := http.ListenAndServe("localhost:8000", nil)

	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}

}

//fetchRSSFeed initiates  action for fetching RSS Feed
func fetchRSSFeed(RSSFeedUrls []string) {

	var errorFromChan error = nil
	rssDataSlice := []rssreader.RSSData{}
	rssDataChan := make(chan []rssreader.RSSData)
	rssErrorChan := make(chan error)

	log.Println("fetching RSS feed...")

	//Fetch the RSS feeds from 2 sources using 2 gouroutines
	for _, url := range RSSFeedUrls {

		url := url

		//we use an anonymous function to make GetRSSFeeds not depend on channels thereby making it more testable
		go func(inputPathOrString string) {

			dataSlice, err := rssreader.GetRSSFeeds(inputPathOrString)

			if err != nil {
				rssErrorChan <- fmt.Errorf("while fetching from %s <= %s", url, err.Error())
				rssDataChan <- []rssreader.RSSData{}
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

		log.Println("[ERROR] main() <= " + errorFromChan.Error())

	}

	storeRSSFeed(rssDataSlice)

}

//storeRSSFeed initiates action to store the RSS Feed in a database
func storeRSSFeed(rssDataSlice []rssreader.RSSData) {

	log.Println("Adding RSS records to database...")

	err := app.DBAdapter.AddRecords(rssDataSlice)

	if err != nil {
		log.Println("[ERROR] main() <= " + err.Error())
	}

}

// This first two if blocks are used for genearting CPU and memory profiles for the application
// We can generate this by running
// `go test -cpuprofile cpu.prof -memprofile mem.prof -bench .` - to generate profiles, and then,
// `go tool pprof --pdf ~/go/src/rss_web_app/cpu.prof > cpu.pdf` - to convert profiles to pdf
func createPerformanceProfile(cpuprofile *string, memprofile *string) {

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

}
