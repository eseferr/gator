package main

import (
	"context"
	"fmt"
	"time"

	"github.com/eseferr/blog-aggregator/internal/database"
)

func handlerAggregator(s *State, cmd Command) error{
	  if len(cmd.Args) < 1 {
        return fmt.Errorf("please provide time between requests (e.g., 1m)")
    }
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
	scrapeFeeds(s)
}
	
}

func scrapeFeeds(s *State)error{
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(),database.MarkFeedFetchedParams{
		ID: nextFeed[0].ID,
		LastFetchedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err !=nil {
		return err
	}
	feed, err := fetchFeed(context.Background(), nextFeed[0].Url)
	if err != nil {
		return err
	}
	for _,item := range feed.Channel.Item{
		fmt.Println(item.Title)
	}
	return nil
}