package handlers

import (
	"errors"
	"fmt"
	"os"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Cmds map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Cmds[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	callback, ok := c.Cmds[cmd.Name]
	if !ok {
		return errors.New("Command does not exist!")
	}
	err := callback(s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return nil
}
