package functions

import (
	"context"
	"fmt"
	"time"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

// add a new feed to the feeds DB. 2 params needed: name and url
func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("Not enough arguments")
	}

	ctx := context.Background()
	var user database.CreateFeedParams

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Name = cmd.Arguments[0]
	user.Url = cmd.Arguments[1]

	currentUser, err := s.Db_ptr.GetUser(ctx, s.Cfg_ptr.CurrentUserName)
	if err != nil {
		return err
	}

	user.UserID = currentUser.ID

	fmt.Printf("%+v", user)

	_, err1 := s.Db_ptr.CreateFeed(ctx, user)
	if err1 != nil {
		return err1
	}

	return nil
}
