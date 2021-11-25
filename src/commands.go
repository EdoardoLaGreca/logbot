package main

import (
	"context"
	"fmt"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

// Routes the commands to their routines, returns an error if there is one.
// If too many arguments are provided, the superfluous will be ignored (e.g. if
// a command accepts 1 argument and 2 are provided, the second will be
// ignored). If not enough arguments are provided, it returns an error.
func cmdRouter(cmd string, args string[]) error {
	// Number of arguments for each command
	nargs := map[string]int {
		"getrpost": 1,
	}

	if len(args) < nargs[cmd] {
		return fmt.Errorf("Not enough arguments for command `" + cmd + "'")
	}


	// Route the commands
	switch (cmd) {
	case "getrpost":
		getRedditPost(args[0], args[1])
	default:
		return fmt.Errorf("Unknown command `" + cmd + "'")
	}

	return nil
}

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

