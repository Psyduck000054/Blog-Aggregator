package functions

import (
	"context"
	"fmt"
	"time"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd Command) error {
	// ensure there is a name in the input
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("No username")
	}

	ctx := context.Background()

	var user database.CreateUserParams

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Name = cmd.Arguments[0]

	createdUser, err := s.Db_ptr.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	err = s.Cfg_ptr.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("user created")
	fmt.Printf("%+v\n", createdUser)

	return nil
}
