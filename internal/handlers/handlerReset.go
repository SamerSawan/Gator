package handlers

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.Db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to reset database: %w", err)
	}
	return nil
}
