package functions

import (
	"context"
	"fmt"
)

func HandlerGetUsers(s *State, cmd Command) error {
	ctx := context.Background()
	db, err := s.Db_ptr.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range db {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.Cfg_ptr.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Print("\n")
	}

	return nil
}
