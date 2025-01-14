package main

import (
	"context"
	"fmt"

	"github.com/eseferr/blog-aggregator/internal/database"
)

func handlerFeedUnfollow(s *State, cmd Command, user database.User) error {
	
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Feed URL is required")
	}
	feedURL := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(),feedURL)
	if err != nil {
		return err
	}
	s.db.FeedUnFollow(context.Background(), database.FeedUnFollowParams{UserID : user.ID, 
		FeedID : feed.ID})
	return nil
}
