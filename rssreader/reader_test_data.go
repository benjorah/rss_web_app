package rssreader

import (
	"time"

	"github.com/mmcdole/gofeed"
)

func getSampleRSSData() (sampleRSSData RSSData) {

	postTime, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2020-02-24 22:10:15 +0000 UTC")
	sampleRSSData = RSSData{
		"Harvey Weinstein found guilty of rape in watershed case",
		"The ex-movie mogul is handcuffed and led from court as a judge orders him to jail immediately.",
		"https://www.bbc.co.uk/news/world-us-canada-51621041",
		&postTime,
	}
	return sampleRSSData
}

func getSampleRSSDataSlice() (sampleRSSDataSlice []RSSData) {

	postTime1, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2020-02-24 22:10:15 +0000 UTC")
	postTime2, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2020-02-24 18:17:58 +0000 UTC")

	sampleRSSDataSlice = []RSSData{
		RSSData{
			"Harvey Weinstein found guilty of rape in watershed case",
			"The ex-movie mogul is handcuffed and led from court as a judge orders him to jail immediately.",
			"https://www.bbc.co.uk/news/world-us-canada-51621041",
			&postTime1,
		},
		RSSData{
			"Harvey Weinstein guilty: How the Hollywood giant faced his reckoning",
			"The judge said this was not a referendum on #MeToo. But at times, the trial felt like one.",
			"https://www.bbc.co.uk/news/world-us-canada-51451977",
			&postTime2,
		},
	}
	return sampleRSSDataSlice

}

func getSampleGoFeedItem() (ampleGoFeedItem *gofeed.Item) {

	postTime, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2020-02-24 22:10:15 +0000 UTC")
	ampleGoFeedItem = &gofeed.Item{
		Title:           "Harvey Weinstein found guilty of rape in watershed case",
		Description:     "The ex-movie mogul is handcuffed and led from court as a judge orders him to jail immediately.",
		Content:         "",
		Link:            "https://www.bbc.co.uk/news/world-us-canada-51621041",
		Updated:         "",
		UpdatedParsed:   nil,
		Published:       "",
		PublishedParsed: &postTime,
	}
	return ampleGoFeedItem
}

const sampleFeed = `
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet title="XSL_formatting" type="text/xsl" href="/shared/bsp/xsl/rss/nolsol.xsl"?>
<rss xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:media="http://search.yahoo.com/mrss/">
    <channel>
        <title><![CDATA[BBC News - World]]></title>
        <description><![CDATA[BBC News - World]]></description>
        <link>https://www.bbc.co.uk/news/</link>
        <image>
            <url>https://news.bbcimg.co.uk/nol/shared/img/bbc_news_120x60.gif</url>
            <title>BBC News - World</title>
            <link>https://www.bbc.co.uk/news/</link>
        </image>
        <generator>RSS for Node</generator>
        <lastBuildDate>Mon, 24 Feb 2020 22:10:44 GMT</lastBuildDate>
        <copyright><![CDATA[Copyright: (C) British Broadcasting Corporation, see http://news.bbc.co.uk/2/hi/help/rss/4498287.stm for terms and conditions of reuse.]]></copyright>
        <language><![CDATA[en-gb]]></language>
        <ttl>15</ttl>
        <item>
            <title><![CDATA[Harvey Weinstein found guilty of rape in watershed case]]></title>
            <description><![CDATA[The ex-movie mogul is handcuffed and led from court as a judge orders him to jail immediately.]]></description>
            <link>https://www.bbc.co.uk/news/world-us-canada-51621041</link>
            <guid isPermaLink="true">https://www.bbc.co.uk/news/world-us-canada-51621041</guid>
            <pubDate>Mon, 24 Feb 2020 22:10:15 GMT</pubDate>
        </item>
        <item>
            <title><![CDATA[Harvey Weinstein guilty: How the Hollywood giant faced his reckoning]]></title>
            <description><![CDATA[The judge said this was not a referendum on #MeToo. But at times, the trial felt like one.]]></description>
            <link>https://www.bbc.co.uk/news/world-us-canada-51451977</link>
            <guid isPermaLink="true">https://www.bbc.co.uk/news/world-us-canada-51451977</guid>
            <pubDate>Mon, 24 Feb 2020 18:17:58 GMT</pubDate>
        </item>
       
    </channel>
</rss>
`
