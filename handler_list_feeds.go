package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	ctx := context.Background()

	db, err := s.db_ptr.GetFeeds(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Feed List:")
	for idx, feed := range db {
		fmt.Printf("%d. \n", idx+1)
		fmt.Printf("	Feed Name: %s\n", feed.FeedName)
		fmt.Printf("	Feed URL : %s\n", feed.Url)
		fmt.Printf("	Creator  : %s\n", feed.Username)
	}

	return nil
}
