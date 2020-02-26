package main

import (
	"rss_web_app/rssreader"
	"time"
)

func getSampleRSSDataSliceForHandler() (sampleRSSDataSlice []rssreader.RSSData) {

	postTime1, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2020-02-24 22:10:15 +0000 UTC")

	sampleRSSDataSlice = []rssreader.RSSData{
		rssreader.RSSData{
			"title",
			"desc",
			"link",
			&postTime1,
		},
	}
	return sampleRSSDataSlice

}
