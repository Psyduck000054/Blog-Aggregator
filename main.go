package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

	var s state
	s.cfg_ptr = &cfg

	// database connection
	db, err := sql.Open("postgres", cfg.DB_URL)
	dbQueries := database.New(db)

	s.db_ptr = dbQueries

	var c commands
	c.Map = make(map[string]func(*state, command) error)

	// ---------------------------------------------------------
	// HANDLERS
	// ---------------------------------------------------------

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerGetUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", handlerAddFeed)
	c.register("feeds", handlerListFeeds)

	if len(os.Args) < 2 {
		fmt.Print(fmt.Errorf("no argument\n"))
		os.Exit(1)
	} else {
		commandName := os.Args[1]
		commandArgs := os.Args[2:]

		cmd := command{
			Name:      commandName,
			Arguments: commandArgs,
		}

		err := c.run(&s, cmd)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
