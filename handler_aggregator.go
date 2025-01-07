package main

import (
	"context"
	"fmt"
)

func handlerAggregator(s *State, cmd Command) error{
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	
	return nil

}