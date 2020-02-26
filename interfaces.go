package main

import "rss_web_app/rssreader"

//DatabaseAdapter is a generic type all database adapters should implement.
//It provides a consistent interface for performing database operation regardless of the database itself
//It also helps in mocking databases
type DatabaseAdapter interface {
	AddRecords(records []rssreader.RSSData) (err error)
	SearchRecords(searchString string) (records []rssreader.RSSData, err error)
}
