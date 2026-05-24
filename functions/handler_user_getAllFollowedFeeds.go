package functions

import (
	"context"
	"fmt"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
)

// input  | 0 param
// output | the name of all feeds that the current user follows
func HandlerGetAllFollowedFeeds(s *State, cmd Command, currentUser database.User) error {
	ctx := context.Background()

	currentUser, err := s.Db_ptr.GetUser(ctx, s.Cfg_ptr.CurrentUserName)
	if err != nil {
		return err
	}

	db, err := s.Db_ptr.GetFFsFromUserID(ctx, currentUser.ID)
	if err != nil {
		return err
	}

	fmt.Printf("The feeds that user %s is following:\n", currentUser.Name)

	for idx, ff := range db {
		feed, err := s.Db_ptr.GetFeedFromID(ctx, ff.FeedID)
		if err != nil {
			return err
		}
		fmt.Printf("	%d. %s\n", idx+1, feed.Name)
	}

	return nil
}
