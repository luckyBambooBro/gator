package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, c command) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	feed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("could not fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}