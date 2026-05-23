package functions

import (
	"fmt"
)

func (c *Commands) Run(s *State, cmd Command) error {
	value, ok := c.Map[cmd.Name]
	if !ok {
		return fmt.Errorf("no command found")
	}

	err := value(s, cmd)

	return err
}

func (c *Commands) Register(name string, f func(s *State, cmd Command) error) {
	c.Map[name] = f
}
