package main

import (
	"reflect"
	"testing"

	"github.com/mmcdole/gofeed"
)

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

func TestReadAndParse(t *testing.T) {
	type args struct {
		inputPathOrString string
	}
	tests := []struct {
		name        string
		args        args
		wantRssFeed []*gofeed.Item
		wantErr     bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRssFeed, err := ReadAndParse(tt.args.inputPathOrString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAndParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRssFeed, tt.wantRssFeed) {
				t.Errorf("ReadAndParse() = %v, want %v", gotRssFeed, tt.wantRssFeed)
			}
		})
	}
}

func BenchmarkReadAndParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadAndParse(sampleFeed)
	}
}
