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

	//check there is a user logged in
	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("you must be logged in to \"add\" feeds")
	}

	//obtain current user
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to retrieve user from users table: %w", err)
	}

	//create params for CreateFeed() below
	feedParams := database.CreateFeedParams {
	ID:        uuid.New(),
	CreatedAt: time.Now().UTC(),
	UpdatedAt: time.Now().UTC(),
	Name:      name,
	Url:       url,
	UserID:    currentUser.ID,
	}

	//add feed to feeds table of database
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("unable to create feed: %w", err)
	}
	fmt.Printf("Added feed:\n%v\n", feed)
	return err
}

func handlerFeeds( s *state, c command) error {
	//get feeds
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to obtain feeds from table: %w", err)
	} 
	if len(feeds) == 0 {
		fmt.Println("No feeds in found in database")
	}
	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.NameFromUsers)
		fmt.Println()
	}
	return nil
}