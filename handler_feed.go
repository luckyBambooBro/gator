package main

import (
	"context"
	"fmt"
	"time"

	"github.com/luckyBambooBro/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, c command) error {
	if len(c.Args) < 2 {
		return fmt.Errorf("not enough arguments provided for \"add\" command")
	}
	name := c.Args[0]
	url := c.Args[1]

	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("you must be logged in to \"add\" feeds")
	}
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to retrieve user from users table: %w", err)
	}
	feedParams := database.CreateFeedParams {
	ID:        uuid.New(),
	CreatedAt: time.Now().UTC(),
	UpdatedAt: time.Now().UTC(),
	Name:      name,
	Url:       url,
	UserID:    currentUser.ID,
	}

	_, err = s.db.CreateFeed(context.Background(), feedParams)
	return err
}