package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/luckyBambooBro/gator/internal/database"
)

func handlerAddFeed(s *state, c command) error {
	if len(c.Args) != 2 {
		return fmt.Errorf("please provide name of feed followed by feed URL")
	}
	name := c.Args[0]
	url := c.Args[1]

	//check there is a user logged in
	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("you must be logged in to \"add\" feeds")
	}

	//obtain current user
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	currentUser, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to retrieve user from users table: %w", err)
	}

	//add feed to feeds table of database
	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	})

	if err != nil {
		return fmt.Errorf("unable to create feed: %w", err)
	}
	fmt.Printf("Feed added successfully:\n%v\n", feed)
	fmt.Println("=====================================")
	return err
}

func handlerListFeeds(s *state, c command) error {
	//get feeds
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	feeds, err := s.db.GetFeeds(ctx)
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

func handlerFollowing (s *state, c command) error {
	if len(c.Args) != 0 {
		return fmt.Errorf("no arguments required for 'following' command")
	}
	//get details of current user
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	userName := s.cfg.CurrentUserName
	userDetails, err := s.db.GetUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("error obtaining user details: %w", err)
	}
	userID := userDetails.ID

	//obtain slice of feedFollows for current user
	feedFollows, err := s.db.GetFeedFollowsForUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("unable to obtain feedFollows for %s", userName)
	}
	//return early if no feeds followed
	if len(feedFollows) == 0 {
		fmt.Printf("No feeds followed for user: %s", userName)
		return nil
	}

	//print name of user and feedFollows for user 
	fmt.Printf("Printing feed follows for %s...\n", userName)
	for _, feedFollow := range feedFollows {
		//obtain the feed using details of feedFollow many to many chart
		feedID := feedFollow.ID
		feed, err := s.db.GetFeedByID(ctx, feedID)
		if err != nil {
			return fmt.Errorf("error obtaining feed: %w", err)
		}
		//print feed name
		fmt.Println(feed.Name)
	}
	return nil
}