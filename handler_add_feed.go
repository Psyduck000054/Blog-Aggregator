package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	// 2 arguments needed: name and url
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

	currentUser, err := s.db_ptr.GetUser(ctx, s.cfg_ptr.CurrentUserName)
	if err != nil {
		return err
	}

	user.UserID = currentUser.ID

	fmt.Printf("%+v", user)

	return nil
}
