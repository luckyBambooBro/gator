package main

import (
	"context"
	"fmt"
	"time"

	"github.com/luckyBambooBro/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	//this function can only take one argument
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <'x's> (x = seconds)", cmd.Name)
	}
	
	//this code block can easily be broken depending on what user types
	selectedTimeBetweenReqs := cmd.Args[0]
	timeBetweenReqs, err := time.ParseDuration(selectedTimeBetweenReqs)
	fmt.Printf("Collecting feeds every %s\n", timeBetweenReqs)
	/*UP TO HERE: Use a time.Ticker to run your scrapeFeeds function once every time_between_reqs. I used a for loop 
	to ensure that it runs immediately (I don't like waiting) and then every time the ticker ticks:*/


	//fetch actual feed 
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	feed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func scrapeFeeds(s *state, cmd command) error {
	//get next feed
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error obtaining feed: %w", err)
	}

	//mark feed as updated
	err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID: feed.ID,
	})

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("unable to obtain rss feed: %w", err)
	}
	for _, rssFeedItem := range rssFeed.Channel.Item {
		fmt.Println(rssFeedItem.Title)
	}

}