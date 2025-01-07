package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)


type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}


func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
	client := http.Client{}
	req, err := http.NewRequestWithContext(ctx,"GET",feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent","gator")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err :=io.ReadAll(resp.Body)
	rss := RSSFeed{}
	err = xml.Unmarshal(respBody,&rss)
	if err != nil {
		return nil,err
	}
	
	// Channel Title and Description
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)

	// Iterate over each RSSItem
	for i := range rss.Channel.Item {
		rss.Channel.Item[i].Title = html.UnescapeString(rss.Channel.Item[i].Title)
		rss.Channel.Item[i].Description = html.UnescapeString(rss.Channel.Item[i].Description)
	}
	return &rss, nil

	}