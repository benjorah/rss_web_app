package database

import (
	"fmt"
	"rss_web_app/rssreader"
	"rss_web_app/system"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

//AlgoliaConnection struct defines methods and data properties for connecting with the algolia databasae
type AlgoliaConnection struct {
	client *search.Client
	index  *search.Index
}

//AddRecords adds new entries into the algolia database
func (algolia *AlgoliaConnection) AddRecords(records []rssreader.RSSData) (err error) {

	_, err = algolia.index.SaveObjects(records, opt.AutoGenerateObjectIDIfNotExist(true))

	return err

}

//SearchRecords searches for entries in the database that matches a particular search string and returns those entries
//It returns an error if any error is encountered
func (algolia *AlgoliaConnection) SearchRecords(searchString string) (records []rssreader.RSSData, err error) {

	res, err := algolia.index.Search(searchString)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] AlgoliaConnection.SearchRecords() : %s", err.Error())
	}

	var RSSDataSlice []rssreader.RSSData

	//Puts the result into a slice of RSSData types
	err = res.UnmarshalHits(&RSSDataSlice)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] AlgoliaConnection.SearchRecords() : %s", err.Error())
	}

	return RSSDataSlice, nil

}

//InitDatabseConnection initializes the algolia client
func (algolia *AlgoliaConnection) InitDatabseConnection() {

	algolia.client = search.NewClient(system.GetEnvVariable("APPLICATION_ID"), system.GetEnvVariable("API_KEY"))
	algolia.index = algolia.client.InitIndex("feed")
}
