package main

import (
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No username")
		os.Exit(1)
	}

	s.ConfigPointer.SetUser(cmd.Arguments[0])

	fmt.Println("user has been set")

	return nil
}
