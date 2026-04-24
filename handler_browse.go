package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/luckyBambooBro/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, u database.User) error {
	var limit = 2
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <limit>", cmd.Name) //check is this the right way to write the error
	} else if len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("unable to obtain limit for %s command", cmd.Name)
		}
	} 

	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	posts, err := s.db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: u.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("unable to obtain posts: %w", err)
	}
	//print posts
	for _, post := range posts {
		fmt.Println("Title: ", post.Title)
		fmt.Println(post.Url)
		if post.Description.Valid {
			fmt.Println(post.Description.String)
		} 
		fmt.Println("Published: ", post.PublishedAt.Format("Jan 02, 2006"))
		fmt.Println("========================")
	}
	return nil
}