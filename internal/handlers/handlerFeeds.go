package handlers

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting feeds: %w", err)
	}
	for _, v := range feeds {
		user, err := s.Db.GetUserByID(context.Background(), v.UserID)
		if err != nil {
			return fmt.Errorf("Error getitng user: %w", err)
		}
		fmt.Println(v.Name)
		fmt.Println(v.Url)
		fmt.Println(user.Name)
	}
	return nil
}
