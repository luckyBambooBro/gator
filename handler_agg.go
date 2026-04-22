package main

import (
	"context"
	"fmt"
	"log"
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
		return fmt.Errorf("invalid duration: %w", err)
	}
	if interval < requestLimit {
		return fmt.Errorf("duration must be at least %v to prevent server overload", requestLimit)
	}
	fmt.Printf("Collecting feeds every %s\n", interval)
	/*UP TO HERE: Use a time.Ticker to run your scrapeFeeds function once every time_between_reqs. I used a for loop 
	to ensure that it runs immediately (I don't like waiting) and then every time the ticker ticks:*/

	ticker := time.NewTicker(interval)
	//doing it this way receives the ticker immediately instead of waiting for the first interval
	for ; ; <-ticker.C {
		scrapeFeeds(s)
		}
}


func scrapeFeeds(s *state) {
	//get next feed
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		log.Printf("error obtaining feed: %v", err)
		return
	}
	scrapeFeed(s.db, feed, ctx)
}

func scrapeFeed(db *database.Queries, feed database.Feed, ctx context.Context) error {
	fmt.Printf("Fetching feeds for: %s\n", feed.Name)

	//mark feed as updated
	err := db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID: feed.ID,
	})
	if err != nil {
		log.Printf("err marking %v feed: %v", feed.Name, err)
		return nil
	}

	rssFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		log.Printf("unable to obtain rss feed: %v", err)
		return nil
	}
	for _, rssFeedItem := range rssFeed.Channel.Item {
		fmt.Println(rssFeedItem.Title)
	}

	fmt.Println()
	fmt.Printf("Feed %s collected, %v posts found\n", feed.Name, len(rssFeed.Channel.Item))
	fmt.Println("================")
	fmt.Println()
	return nil
	}
		

