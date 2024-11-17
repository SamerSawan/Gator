package handlers

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("You must enter a username!")
	}
	_, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println("User does not exist!")
		os.Exit(1)
	}
	err = s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", cmd.Args[0])
	return nil
}
