package handlers

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("You must enter a username!")
	}
	err := s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", cmd.Args[0])
	return nil
}
