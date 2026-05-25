package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	feed := &RSSFeed{}
	var reader io.Reader

	// http get the xml data in the url
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, reader)
	if err != nil {
		return feed, err
	}

	// http client do
	cli := &http.Client{}
	res0, err := cli.Do(req)
	if err != nil {
		return feed, err
	}

	// readall
	res1, err := io.ReadAll(res0.Body)

	// unmarshal
	err = xml.Unmarshal(res1, feed)
	if err != nil {
		return feed, err
	}

	// decode escaped html entities in title and desc
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for index, item := range feed.Channel.Item {
		feed.Channel.Item[index].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[index].Description = html.UnescapeString(item.Description)
	}

	return feed, nil
}
