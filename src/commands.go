package main

import (
	"context"
	"fmt"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

// Get a Reddit post
// Receives the name of a subreddit and returns the URL of a post from that
// subreddit. Returns an error if the subreddit has no new posts or another
// error happened
func getRedditPost(sub string) (string, error) {
	client := reddit.DefaultClient()

	// Get new posts
	posts, _, err := client.Subreddit.NewPosts(context.Background(), sub, nil)
	if err != nil {
		return "", err
	}

	if len(posts) > 0 {
		return posts[0].URL
	} else {
		return "", fmt.Errorf("No new posts")
	}
}
