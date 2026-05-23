package functions

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	ctx := context.Background()
	err := s.Db_ptr.DeleteUsers(ctx)
	if err == nil {
		fmt.Println("reset success!")
	} else {
		fmt.Println("reset failed!")
	}
	return err
}
