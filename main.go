package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Psyduck000054/Blog-Aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	var s state
	s.ConfigPointer = &cfg

	var c commands
	c.Map = make(map[string]func(*state, command) error)

	c.register("login", handlerLogin)

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
			fmt.Print(fmt.Errorf("failed run\n"))
			os.Exit(1)
		}
	}
}
