package main

import (
	"context"
	"fmt"
	"time"

	"github.com/luckyBambooBro/gator/internal/database"
)

const requestLimit = 3 * time.Second

func handlerAgg(s *state, cmd command) error {
	//this function can only take one argument
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <duration>", cmd.Name)
	}
	
	//this code block can easily be broken depending on what user types
	interval, err := time.ParseDuration(cmd.Args[0]) //i just used cmd.Args[0] but the lesson wanted me to label it as 
	//a variable called time_between_reqs. i dont think it matters?
	if err != nil {
		return fmt.Errorf("invaliud duration: %w", err)
	}
	if interval < requestLimit {
		return fmt.Errorf("duration must be at least 1s to prevent server overload")
	}
	fmt.Printf("Collecting feeds every %s\n", interval)
	/*UP TO HERE: Use a time.Ticker to run your scrapeFeeds function once every time_between_reqs. I used a for loop 
	to ensure that it runs immediately (I don't like waiting) and then every time the ticker ticks:*/

	ticker := time.NewTicker(interval)
	//doing it this way receives the ticker immediately instead of waiting for the first interval
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Printf("error scraping feed: %v", err)
			continue
		}
		}

}

func scrapeFeeds(s *state) error {
	//get next feed
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error obtaining feed: %w", err)
	}
	err = scrapeFeed(s.db, feed, ctx)
	if err != nil {
		fmt.Printf("error scraping feed: %v", err)
	}
	return nil
}

func scrapeFeed(db *database.Queries, feed database.Feed, ctx context.Context) error {
	fmt.Printf("Fetching feeds for: %s\n", feed.Name)

	//mark feed as updated
	err := db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
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
	fmt.Println("================")
	fmt.Println()
	return nil
	}
		

