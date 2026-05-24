package functions

import (
	"context"
	"fmt"
)

// input  | 0 param
// output | state all users in users DB and mark the current one
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
