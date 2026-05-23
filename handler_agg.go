package main

import (
	"context"
	"fmt"

	"github.com/Psyduck000054/Blog-Aggregator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	link := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	feed, err := rss.FetchFeed(ctx, link)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", *feed)

	return nil
}
