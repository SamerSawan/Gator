package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/samersawan/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit argument: %v", err)
		}
		limit = parsedLimit
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limit)})
	if err != nil {
		return fmt.Errorf("Failed to get posts: %w", err)
	}
	for _, v := range posts {
		fmt.Println(v.Title)
		fmt.Println(v.Description)
		fmt.Println("------------------------------")
	}
	return nil
}
