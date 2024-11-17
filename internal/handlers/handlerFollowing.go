package handlers

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("Failed to fetch user: %w", err)
	}
	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Failed to fetch feeds: %w", err)
	}
	for _, v := range feeds {
		fmt.Println(v.FeedName)
	}
	return nil
}
