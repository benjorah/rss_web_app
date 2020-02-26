package main

import "time"

func getSampleRSSDataSliceForHandler() (sampleRSSDataSlice []RSSData) {

	postTime1, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2020-02-24 22:10:15 +0000 UTC")

	sampleRSSDataSlice = []RSSData{
		RSSData{
			"title",
			"desc",
			"link",
			&postTime1,
		},
	}
	return sampleRSSDataSlice

}
