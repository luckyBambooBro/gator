package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/luckyBambooBro/gator/internal/database"
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
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%q does not exist in database. please add feed to database to follow", url)
		} 
		return fmt.Errorf("error following %q: %w", url, err )
		
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
	//the following error check was advice i learnt from gemini. itll give a more 
	//specific error, rather than just returning err. even though we checked the feed doesnt exist above
	//apparently it is still good to have this check
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return fmt.Errorf("you are already following this feed")
			}
		}
		return err
	}

	//use the feedFollow to print user name and feed name 
	fmt.Printf("Created feed %v for user: %v\n", feedFollowRecord.FeedName, feedFollowRecord.UserName)
	return nil
}

//prints all the feeds being followed by the current user
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

	//obtain feeds followed by user
	feedsFollowed, err := s.db.GetFeedFollowsForUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("error obtaining feeds for user: %s\n", userName)
	}

	if len(feedsFollowed) == 0 {
		fmt.Printf("no feeds followed by user: %s\n", userName)
		return nil
	}

	//print feeds followed
	fmt.Printf("Printing feeds for %s...\n", userName)
	for _, feedFollowed := range feedsFollowed {
		fmt.Println("- ", feedFollowed.FeedName)
	}
	return nil

}