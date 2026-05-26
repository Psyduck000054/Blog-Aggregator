package functions

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
)

// input  | 1 param: limit [optional]
// output | return [limit] first posts that the users followed
func HandlerBrowse(s *State, cmd Command, currentUser database.User) error {
	var limit int
	var err0 error
	if len(cmd.Arguments) > 0 {
		limit, err0 = strconv.Atoi(cmd.Arguments[0])
		if err0 != nil {
			return err0
		}
	} else {
		limit = 2
	}

	lim32 := int32(limit)

	ctx := context.Background()

	input0 := database.GetPostsForUserParams{
		UserID: currentUser.ID,
		Limit:  lim32,
	}

	postList, err1 := s.Db_ptr.GetPostsForUser(ctx, input0)
	if err1 != nil {
		return err1
	}

	fmt.Printf("First %d posts that user %s follows:\n", limit, currentUser.Name)
	for index, item := range postList {
		var nulltitle string
		if item.Title == "" {
			nulltitle = "No title"
		} else {
			nulltitle = item.Title
		}
		fmt.Printf("	%d. %s\n", index+1, nulltitle)
	}

	return nil
}
