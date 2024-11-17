package handlers

import (
	"context"
	"fmt"
)

func HandlerGetUsers(s *State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching users: %w", err)
	}
	for _, v := range users {
		if v.Name == s.Cfg.CurrentUsername {
			fmt.Printf("* %s (current)\n", v.Name)
		} else {
			fmt.Printf("* %s\n", v.Name)
		}
	}
	return nil
}
