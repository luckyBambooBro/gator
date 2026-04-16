package main

import (
	"context"
	"fmt"

	"github.com/luckyBambooBro/gator/internal/database"
)

func middlewareLoggedIn(handler func(*state, command, database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// get current user
		ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
		defer cancel()
		user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("unable to obtain current user: %w", err)
		}
		// run the handler (s and cmd will come from as they are passed into cmds.run() in main.go, user comes from this function)
		if err = handler(s, cmd, user); err != nil {
			return fmt.Errorf("error running %q command: %w", cmd.Name, err)
		}
		return nil 
	}
}
