package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No username")
	}

	ctx := context.Background()

	_, err := s.db_ptr.GetUser(ctx, cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("No user named like this in database")
	}

	s.cfg_ptr.SetUser(cmd.Arguments[0])

	fmt.Println("user has been set")

	return nil
}
