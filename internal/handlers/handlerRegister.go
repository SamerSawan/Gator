package handlers

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/samersawan/gator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("You must enter a username!")
	}
	userParams := database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.Args[0]}
	user, err := s.Db.CreateUser(context.Background(), userParams)
	if err != nil {
		fmt.Println("User already exists!")
		os.Exit(1)
	}
	s.Cfg.SetUser(cmd.Args[0])
	fmt.Println(user)
	return nil
}
