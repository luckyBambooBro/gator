package main

import (
	"context"
	"fmt"

	"github.com/luckyBambooBro/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// get current user
		ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
		defer cancel()
		currentUser, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("unable to obtain current user: %w", err)
		}
	//up to here
		
		return nil 
	}
}
