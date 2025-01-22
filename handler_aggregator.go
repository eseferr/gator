package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/eseferr/blog-aggregator/internal/database"
	"github.com/google/uuid"
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
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time: t,
				Valid: true,
			}
		}
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    nextFeed[0].ID,
			Title:     sql.NullString{
				String: item.Title,
				Valid: true,
			},
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publishedAt,

		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", nextFeed[0].Name, len(feed.Channel.Item))
	return nil
	}
	