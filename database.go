package main

import (
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

//AlgoliaConnection struct defines methods and data properties for connecting with the algolia databasae
type AlgoliaConnection struct {
	client *search.Client
	index  *search.Index
}

//AddRecords adds new entries into the algolia database
func (algolia *AlgoliaConnection) AddRecords(records []RSSData) (err error) {

	_, err = algolia.index.SaveObjects(records, opt.AutoGenerateObjectIDIfNotExist(true))

	return err

}

//SearchRecords searches for entries in the database that matches a particular search string and returns those entries
//It returns an error if any error is encountered
func (algolia *AlgoliaConnection) SearchRecords(searchString string) (records []RSSData, err error) {

	res, err := algolia.index.Search(searchString)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] AlgoliaConnection.SearchRecords() : %s", err.Error())
	}

	var RSSDataSlice []RSSData

	//Puts the result into a slice of RSSData types
	err = res.UnmarshalHits(&RSSDataSlice)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] AlgoliaConnection.SearchRecords() : %s", err.Error())
	}

	return RSSDataSlice, nil

}

//InitDatabseConnection initializes the algolia client
func (algolia *AlgoliaConnection) InitDatabseConnection() {

	algolia.client = search.NewClient(goDotEnvVariable("APPLICATION_ID"), goDotEnvVariable("API_KEY"))
	algolia.index = algolia.client.InitIndex("feed")
}
