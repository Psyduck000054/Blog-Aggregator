package functions

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
	"github.com/Psyduck000054/Blog-Aggregator/internal/rss"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

		//null desc
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		// null time
		var publishedTime sql.NullTime
		pubDate, err := time.Parse(time.RFC1123, item.PubDate)
		if err == nil {
			publishedTime.Time = pubDate
			publishedTime.Valid = true
		} else {
			pubDate, err = time.Parse(time.RFC1123Z, item.PubDate)
			if err == nil {
				publishedTime.Time = pubDate
				publishedTime.Valid = true
			} else {
				publishedTime.Valid = false
			}
		}

		newPost := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedTime,
			FeedID:      feed.ID,
		}

		_, err1 := s.Db_ptr.CreatePost(ctx, newPost)
		if err1 != nil {
			pqErr, ok := err1.(*pq.Error)
			if ok {
				if string(pqErr.Code) == "23505" {
					continue
				} else {
					fmt.Printf("Error: %v\n", err1)
					return err1
				}
			} else {
				fmt.Printf("Error: %v\n", err1)
				return err1
			}
		}
	}
	return nil
}
