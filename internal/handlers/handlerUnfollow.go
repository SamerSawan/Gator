package handlers

import (
	"context"
	"fmt"

	"github.com/samersawan/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Incorrect usage! Expected %s <url>", cmd.Name)
	}
	err := s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("Failed to delete FeedFollow connection: %w", err)
	}
	return nil
}
