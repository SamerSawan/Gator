package handlers

import (
	"context"
	"fmt"

	"github.com/samersawan/gator/internal/config"
	"github.com/samersawan/gator/internal/database"
)

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUsername)
		if err != nil {
			return fmt.Errorf("failed to fetch user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
