package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

const (
	timeout = 5 * time.Second
	userAgent = "gator"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	//start new request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	//set header to allow websites to identify us
	req.Header.Set("User-Agent", userAgent)

	//obtain and close http.Response.Body
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return  nil, err
	}
	defer resp.Body.Close()

	//readall
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//data is a []bytes that we need to unmarshal using xml.Unmarshal()
	rss := &RSSFeed{}
	if err = xml.Unmarshal(data, rss); err != nil {
		return nil, err
	}

	/*Use the html.UnescapeString function to decode escaped HTML entities (like &ldquo;). You'll need to run the Title and 
Description fields (of both the entire channel as well as the items) through this function.*/

	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
	for i := range rss.Channel.Item {
		rss.Channel.Item[i].Title = html.UnescapeString(rss.Channel.Item[i].Title) 
		rss.Channel.Item[i].Description = html.UnescapeString(rss.Channel.Item[i].Description)
	}
	return rss, nil
}




