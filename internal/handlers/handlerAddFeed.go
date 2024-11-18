package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samersawan/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Incorrect usage. Expected %s <name> <url>", cmd.Name)
	}
	newFeed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{Name: cmd.Args[0], Url: cmd.Args[1], UserID: user.ID})
	if err != nil {
		return fmt.Errorf("Failed to create feed: %w", err)
	}
	fmt.Println(newFeed)
	s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    newFeed.Url,
	})
	return nil
}
