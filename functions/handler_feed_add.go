package functions

import (
	"context"
	"fmt"
	"time"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

// input  | 2 param: feed name and feed url
// output | add the feed as a datarow in feeds DB
func HandlerAddFeed(s *State, cmd Command, currentUser database.User) error {
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("Not enough arguments, need 2: [name] and [url]")
	}

	ctx := context.Background()
	var feed database.CreateFeedParams

	feed.ID = uuid.New()
	feed.CreatedAt = time.Now()
	feed.UpdatedAt = time.Now()
	feed.Name = cmd.Arguments[0]
	feed.Url = cmd.Arguments[1]

	feed.UserID = currentUser.ID

	fmt.Printf("%+v\n", feed)

	createdFeed, err1 := s.Db_ptr.CreateFeed(ctx, feed)
	if err1 != nil {
		return err1
	}

	feedCommand := Command{
		Arguments: []string{createdFeed.Url},
	}

	// create a user-feed link as a datarow in FFs DB
	err2 := HandlerAddFollow(s, feedCommand, currentUser)
	if err2 != nil {
		return err2
	}

	return nil
}
