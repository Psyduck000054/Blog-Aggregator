package functions

import (
	"fmt"
)

// input  | 0 param
// output | print all available commands and their usage
func HandlerHelp(s *State, cmd Command) error {
	helpText := `
Welcome to Blog-Aggregator!
Usage:

// ---------------------------------------------------------
// USER COMMANDS
// ---------------------------------------------------------
  login <username>      - Log in as an existing user
  register <username>   - Create a new user account
  users                 - List all users
  reset                 - Reset the database (developer use)

// ---------------------------------------------------------
// FEED COMMANDS
// ---------------------------------------------------------
  addfeed <name> <url>  - Add a new RSS feed and follow it
  feeds                 - List all available feeds
  follow <url>          - Follow an existing feed
  following             - List all feeds you are following
  unfollow <url>        - Unfollow a feed

// ---------------------------------------------------------
// AGGREGATION COMMANDS
// ---------------------------------------------------------
  agg <time_interval>   - Start the background scraper (e.g., 1m, 1h)
  browse [limit]        - Browse the [limit] latest posts from followed feeds.
                          Default limit to 2 if there is no parameter after "browse"
  help                  - Show this help message
`
	fmt.Print(helpText)
	return nil
}
