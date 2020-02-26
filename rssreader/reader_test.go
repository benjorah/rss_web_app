package rssreader

import (
	"reflect"
	"testing"

	"github.com/mmcdole/gofeed"
)

/*

	Tests for ReadAndParse() STARTS here
*/
func TestReadAndParse(t *testing.T) {
	type args struct {
		inputPathOrString string
	}
	tests := []struct {
		name            string
		args            args
		wantRssFeedItem []*gofeed.Item
		wantErr         bool
	}{
		{
			name:            "it should return an error and nil if string is empty",
			args:            args{""},
			wantRssFeedItem: nil,
			wantErr:         true,
		},

		{
			name:            "it should return an error and znil if string is not of RSS format",
			args:            args{"This string is not of RSS format so it fails"},
			wantRssFeedItem: nil,
			wantErr:         true,
		},

		{
			name:            "it should return an error and nil if string is a file path to a non xml file",
			args:            args{"models.go"},
			wantRssFeedItem: nil,
			wantErr:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotRssFeed, err := ReadAndParse(tt.args.inputPathOrString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAndParse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRssFeed, tt.wantRssFeedItem) {
				t.Errorf("ReadAndParse() = %v, want %v", gotRssFeed, tt.wantRssFeedItem)
			}
		})
	}

	//Other tests that won't be convenient as aa table test
	sampleRSSData := getSampleRSSData()
	t.Run("It should return the RSSData if the input string contains RSS data", func(t *testing.T) {
		t.Parallel()
		gotRssFeed, err := ReadAndParse(sampleFeed)
		if err != nil {
			t.Errorf("ReadAndParse() error = %v, wantErr %v", err, nil)
			return
		}
		if !reflect.DeepEqual(gotRssFeed[0].Title, sampleRSSData.Title) {
			t.Errorf("ReadAndParse() = %v, want %v", *gotRssFeed[0], sampleRSSData)
		} else if !reflect.DeepEqual(gotRssFeed[0].Description, sampleRSSData.Description) {
			t.Errorf("ReadAndParse() = %v, want %v", *gotRssFeed[0], sampleRSSData)
		} else if !reflect.DeepEqual(gotRssFeed[0].Link, sampleRSSData.Link) {
			t.Errorf("ReadAndParse() = %v, want %v", *gotRssFeed[0], sampleRSSData)
		} else if !reflect.DeepEqual(gotRssFeed[0].PublishedParsed, sampleRSSData.CreatedAt) {
			t.Errorf("ReadAndParse() = %v, want %v", *gotRssFeed[0], sampleRSSData)
		}

	})

	t.Run("It should return the RSSData if input string is a file path to an xml file", func(t *testing.T) {
		t.Parallel()
		gotRssFeed, err := ReadAndParse("reader_test_data.xml")
		if err != nil {
			t.Errorf("ReadAndParse() error = %v, wantErr %v", err, nil)
			return
		}
		if !reflect.DeepEqual(gotRssFeed[0].Title, sampleRSSData.Title) {
			t.Errorf("ReadAndParse() = %v, want %v", *gotRssFeed[0], sampleRSSData)
		} else if !reflect.DeepEqual(gotRssFeed[0].Description, sampleRSSData.Description) {
			t.Errorf("ReadAndParse() = %v, want %v", *gotRssFeed[0], sampleRSSData)
		} else if !reflect.DeepEqual(gotRssFeed[0].Link, sampleRSSData.Link) {
			t.Errorf("ReadAndParse() = %v, want %v", *gotRssFeed[0], sampleRSSData)
		} else if !reflect.DeepEqual(gotRssFeed[0].PublishedParsed, sampleRSSData.CreatedAt) {
			t.Errorf("ReadAndParse() = %v, want %v", *gotRssFeed[0], sampleRSSData)
		}

	})

}

func BenchmarkReadAndParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadAndParse(sampleFeed)
	}
}

/*

	Tests for ReadAndParse() ENDS here
*/

/*

	Tests for GetRSSFeeds() STARTS here
*/
func TestGetRSSFeeds(t *testing.T) {
	type args struct {
		inputPathOrString string
	}
	tests := []struct {
		name     string
		args     args
		wantData []RSSData
		wantErr  bool
	}{
		{
			name:     "it should return an error and zero value slice if string is empty",
			args:     args{""},
			wantData: nil,
			wantErr:  true,
		},

		{
			name:     "it should return an error and zero value slice if string is not of RSS format",
			args:     args{"This string is not of RSS format so it fails"},
			wantData: nil,
			wantErr:  true,
		},

		{
			name:     "it should return an error and zero value slice if string is a file path to a non xml file",
			args:     args{"models.go"},
			wantData: nil,
			wantErr:  true,
		},

		{
			name:     "It should return the RSSData slice if input string is a file path to an xml file",
			args:     args{"reader_test_data.xml"},
			wantData: getSampleRSSDataSlice(),
			wantErr:  false,
		},

		{
			name:     "It should return the RSSData slice if the input string contains RSS data",
			args:     args{sampleFeed},
			wantData: getSampleRSSDataSlice(),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := GetRSSFeeds(tt.args.inputPathOrString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRSSFeeds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetRSSFeeds() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func BenchmarkGetRSSFeeds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRSSFeeds(sampleFeed)
	}
}

/*

	Tests for GetRSSFeeds() ENDS here
*/

/*

	Tests for TransFormRSSItemToCustomData() STARTS here
*/
func TestTransFormRSSItemToCustomData(t *testing.T) {
	type args struct {
		feedItem *gofeed.Item
	}
	tests := []struct {
		name     string
		args     args
		wantData RSSData
		wantErr  bool
	}{
		{
			"It should return the transformed data",
			args{getSampleGoFeedItem()},
			getSampleRSSData(),
			false,
		},

		{
			"It should return an error and empty RSSData when the Feed Item is not valid",
			args{nil},
			RSSData{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := TransFormRSSItemToCustomData(tt.args.feedItem)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransFormRSSItemToCustomData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("TransFormRSSItemToCustomData() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func BenchmarkTransFormRSSItemToCustomData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TransFormRSSItemToCustomData(getSampleGoFeedItem())
	}
}

/*

	Tests for TransFormRSSItemToCustomData() ENDS here
*/
