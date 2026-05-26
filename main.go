package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Psyduck000054/Blog-Aggregator/functions"
	"github.com/Psyduck000054/Blog-Aggregator/internal/config"
	"github.com/Psyduck000054/Blog-Aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	// ---------------------------------------------------------
	// READ CONFIG
	// ---------------------------------------------------------
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// ---------------------------------------------------------
	// STATE INITIALIZATION
	// ---------------------------------------------------------

	var s functions.State
	s.Cfg_ptr = &cfg

	// database connection
	db, err := sql.Open("postgres", cfg.DB_URL)
	dbQueries := database.New(db)

	s.Db_ptr = dbQueries

	var c functions.Commands
	c.Map = make(map[string]func(*functions.State, functions.Command) error)

	// ---------------------------------------------------------
	// HANDLERS
	// ---------------------------------------------------------

	c.Register("login", functions.HandlerLogin)
	c.Register("register", functions.HandlerRegister)
	c.Register("reset", functions.HandlerReset)
	c.Register("users", functions.HandlerGetUsers)
	c.Register("agg", functions.HandlerAgg)
	c.Register("addfeed", functions.MiddlewareLoggedIn(functions.HandlerAddFeed))
	c.Register("feeds", functions.HandlerListFeeds)
	c.Register("follow", functions.MiddlewareLoggedIn(functions.HandlerAddFollow))
	c.Register("following", functions.MiddlewareLoggedIn(functions.HandlerGetAllFollowedFeeds))
	c.Register("unfollow", functions.MiddlewareLoggedIn(functions.HandlerDeleteFollowedFeed))
	c.Register("browse", functions.MiddlewareLoggedIn(functions.HandlerBrowse))

	if len(os.Args) < 2 {
		fmt.Print(fmt.Errorf("no argument\n"))
		os.Exit(1)
	} else {
		commandName := os.Args[1]
		commandArgs := os.Args[2:]

		cmd := functions.Command{
			Name:      commandName,
			Arguments: commandArgs,
		}

		// ---------------------------------------------------------
		// EXECUTION
		// ---------------------------------------------------------

		err := c.Run(&s, cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
