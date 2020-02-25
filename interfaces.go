package main

//DatabaseAdapter is a generic type all database adapters should implement.
//It provides a consistent interface for performing database operation regardless of the database itself
//It also helps in mocking databases
type DatabaseAdapter interface {
	AddRecords(records []RSSData) (err error)
	SearchRecords(searchString string) (records []RSSData, err error)
}
