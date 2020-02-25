package main

import "time"

//RSSData defines the internal structural representatation for an individual RSS Feed Item
type RSSData struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Link        string     `json:"link"`
	CreatedAt   *time.Time `json:"created_at"`
}

//HTTPResponse defines the internal structural representatation for a http response object (success)
type HTTPResponse struct {
	Success         bool `json:"success"`
	ResponsePayload `json:"payload"`
	ResponseError   `json:"error"`
}

//ResponsePayload defines the structure of a payload  object contained in a http response
type ResponsePayload struct {
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

//ResponseError defines the structure of an error object contained in a http response
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//App defines the structural representatation for the application.
//It takes a database adapter(as an interface to allow any adapter implementing the interface to satisfy the requiremnet)
//which it uses to perform operation on said databse
type App struct {
	DBAdapter DatabaseAdapter
}
