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

func handlerFollow(s *state, c command, u database.User) error {
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

	//create the feed
	feedFollowRecord, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: u.ID,
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
func handlerListFeedFollows (s *state, c command, u database.User) error {
	if len(c.Args) != 0 {
		return fmt.Errorf("no arguments required for 'following' command")
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	//obtain feeds followed by user
	feedsFollowed, err := s.db.GetFeedFollowsForUser(ctx, u.ID)
	if err != nil {
		return fmt.Errorf("error obtaining feeds for user: %s\n", u.Name)
	}

	if len(feedsFollowed) == 0 {
		fmt.Printf("no feeds followed by user: %s\n", u.Name)
		return nil
	}

	//print feeds followed
	fmt.Printf("Printing feeds for %s...\n", u.Name)
	for _, feedFollowed := range feedsFollowed {
		fmt.Println("- ", feedFollowed.FeedName)
	}
	return nil

}

func handlerUnfollowFeed(s *state, cmd command, u database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_URL>", cmd.Name)
	}
	//obtain feedID
	feedURL := cmd.Args[0]
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	feed, err := s.db.GetFeedByURL(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("unable to obtain feed: %w", err)
	}

	//use unfollow command
	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: u.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to unfollow feed: %q. error: %w", feedURL, err)
	}
	fmt.Printf("Successfully unfollowed %q feed for user %q", feedURL, u.Name)
	return nil
}