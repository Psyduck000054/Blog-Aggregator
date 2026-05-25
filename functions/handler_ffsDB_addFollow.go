package functions

import (
	"context"
	"fmt"
	"time"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

// input  | 1 param: the feed's url
// output | create a new user-feed link as a datarow in FFs DB
func HandlerAddFollow(s *State, cmd Command, currentUser database.User) error {
	ctx := context.Background()

	// no argument: go run . follow
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("need a feed url argument")
	}

	currentFeed, err := s.Db_ptr.GetFeedFromURL(ctx, cmd.Arguments[0])
	if err != nil {
		return err
	}

	feed_follow_0 := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    currentFeed.ID,
	}

	dataRow, err := s.Db_ptr.CreateFeedFollow(ctx, feed_follow_0)
	if err != nil {
		fmt.Println("Follow Request Failed!")
		return err
	}

	fmt.Println("Follow Request Success!")
	fmt.Printf("user %s has followed feed %s\n", dataRow.UserName, dataRow.FeedName)

	return nil
}
