package main

type Databaser interface {
	AddRecords(records []RSSData) (err error)
	SearchRecords(searchString string) (data []RSSData, err error)
}
