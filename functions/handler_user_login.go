package functions

import (
	"context"
	"fmt"
)

// input  | 1 param: username
// output | check if the user is in users DB
//
//	and if yes, make the user the current user
func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No username")
	}

	ctx := context.Background()

	_, err := s.Db_ptr.GetUser(ctx, cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("No user named like this in database")
	}

	s.Cfg_ptr.SetUser(cmd.Arguments[0])

	fmt.Println("user has been set")

	return nil
}
