package main

import (
	"fmt"
)

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.Map[cmd.Name]
	if !ok {
		return fmt.Errorf("no command found")
	}

	err := value(s, cmd)

	return err
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.Map[name] = f
}
