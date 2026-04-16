package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/luckyBambooBro/gator/internal/database"
)

func handlerAddFeed(s *state, c command, u database.User) error {
	if len(c.Args) != 2 {
		return fmt.Errorf("please provide name of feed followed by feed URL")
	}
	name := c.Args[0]
	url := c.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()


	//add feed to feeds table of database
	feed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    u.ID,
	})
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return fmt.Errorf("this feed already exists in the database")
			}
		}
		return fmt.Errorf("unable to create feed: %w", err)
	}
	fmt.Printf("Feed created successfully: %q\n", url)

	//automatically create a feedFollow for the user for the feed we just created
	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: u.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("")
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

