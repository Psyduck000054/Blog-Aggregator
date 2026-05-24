package functions

import (
	"context"
	"fmt"
)

// input  | 0 param
// output | get all feeds from feeds DB
func HandlerListFeeds(s *State, cmd Command) error {
	ctx := context.Background()

	db, err := s.Db_ptr.GetFeeds(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Feed List:")
	for idx, feed := range db {
		fmt.Printf("%d. \n", idx+1)
		fmt.Printf("	Feed Name: %s\n", feed.FeedName)
		fmt.Printf("	Feed URL : %s\n", feed.FeedUrl)
		fmt.Printf("	Creator  : %s\n", feed.Username)
	}

	return nil
}
