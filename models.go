package main

import "time"

//RSSData defines the internal structural representatation for an individual RSS Feed Item
type RSSData struct {
	Title string `json:"title"`

	Description string `json:"description"`

	Link string `json:"link"`

	CreatedAt *time.Time `json:"created_at"`
}

type App struct {
	Database Databaser
}
