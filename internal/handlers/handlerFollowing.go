package handlers

import (
	"context"
	"fmt"

	"github.com/samersawan/gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Failed to fetch feeds: %w", err)
	}
	for _, v := range feeds {
		fmt.Println(v.FeedName)
	}
	return nil
}
