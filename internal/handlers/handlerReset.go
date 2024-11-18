package handlers

import (
	"context"
	"fmt"

	"github.com/samersawan/gator/internal/database"
)

func HandlerReset(s *State, cmd Command, user database.User) error {
	err := s.Db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to reset database: %w", err)
	}
	return nil
}
