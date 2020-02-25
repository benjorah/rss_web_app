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

	//program starts
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

	//Fetch the RSS feeds from 2 sources using 2 gouroutines

	go GetRSSFeeds("http://feeds.bbci.co.uk/news/world/rss.xml", rssDataChan, rssErrorChan1)
	go GetRSSFeeds("http://rss.cnn.com/rss/edition_world.rss", rssDataChan2, rssErrorChan2)

	//Wait for Data and Error
	rssDataSlice = append(rssDataSlice, <-rssDataChan...)

	rssDataSlice = append(rssDataSlice, <-rssDataChan2...)

	erroFromChan1 := <-rssErrorChan1
	erroFromChan2 := <-rssErrorChan2

	if erroFromChan1 != nil {

		log.Println("[ERROR] main() <= " + erroFromChan1.Error())

	}

	if erroFromChan2 != nil {

		log.Println("[ERROR] main() <= " + erroFromChan2.Error())

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
