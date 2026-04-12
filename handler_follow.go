package main

import (
	"context"
	"fmt"
	"time"
)

func handlerFollows(s *state, c command) error {
	if len(c.Args) != 1 {
		return fmt.Errorf("please provide one URL for follow command")
	}

	//obtain feed by looking up URL
	url := c.Args[0]
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	feed, err := s.db.GetFeedByURL(ctx, url)
}