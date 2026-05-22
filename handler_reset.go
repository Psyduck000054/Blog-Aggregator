package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db_ptr.DeleteUsers(ctx)
	if err == nil {
		fmt.Println("reset success!")
	} else {
		fmt.Println("reset failed!")
	}
	return err
}
