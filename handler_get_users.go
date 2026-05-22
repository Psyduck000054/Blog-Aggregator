package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	ctx := context.Background()
	db, err := s.db_ptr.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range db {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.cfg_ptr.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Print("\n")
	}

	return nil
}
