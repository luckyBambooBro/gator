package main

import (
	"context"
	"fmt"
	"time"

	"github.com/luckyBambooBro/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, c command) error {
	if len(c.Args) != 1 {
		return fmt.Errorf("please provide one URL for follow command")
	}

	//obtain feed by looking up URL
	url := c.Args[0]
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("error obtaining feed: %w", err)
	}

	//obtain user details
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	//create the feed
	feedFollowRecord, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	//use the feedFollow to print user name and feed name 
	fmt.Printf("Created feed %v for user: %v\n", feedFollowRecord.FeedName, feedFollowRecord.UserName)
	return nil
}