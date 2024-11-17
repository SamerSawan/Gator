package handlers

import (
	"context"
	"fmt"
	"os"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.Db.Reset(context.Background())
	if err != nil {
		fmt.Println("Failed to reset database")
		os.Exit(1)
	}
	return nil
}
