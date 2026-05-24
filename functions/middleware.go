package functions

import (
	"context"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.Db_ptr.GetUser(context.Background(), s.Cfg_ptr.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
