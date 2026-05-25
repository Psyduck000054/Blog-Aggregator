package functions

import (
	"context"
	"fmt"

	"github.com/Psyduck000054/Blog-Aggregator/internal/rss"
)

// input  | 0 param
// output | get the info of the next feed to fetch, prio null > oldest
// then print every of the items's title
func HandlerScrapeFeeds(s *State, cmd Command) error {
	ctx := context.Background()
	feed, err := s.Db_ptr.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	err = s.Db_ptr.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := rss.FetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("Feed %s's items:\n", rssFeed.Channel.Title)
	for idx, item := range rssFeed.Channel.Item {
		fmt.Printf("	%d. %s\n", idx+1, item.Title)
	}
	return nil
}
