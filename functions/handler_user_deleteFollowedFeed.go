package functions

import (
	"context"
	"fmt"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
)

// input  | 1 param: feed url
// output | if the user is following the feed which its url is the param,
// delete the datarow contains both of them in FFs DB
func HandlerDeleteFollowedFeed(s *State, cmd Command, currentUser database.User) error {
	ctx := context.Background()

	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("not enough arguments")
	}

	feed, err := s.Db_ptr.GetFeedFromURL(ctx, cmd.Arguments[0])
	if err != nil {
		return err
	}

	paramsStruct := database.DeleteFeedFollowParams{
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}

	err = s.Db_ptr.DeleteFeedFollow(ctx, paramsStruct)
	if err != nil {
		return err
	}

	return nil
}
