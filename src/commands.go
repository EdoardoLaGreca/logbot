package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/diamondburned/arikawa/v3/session"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

// Routes the commands to their routines, handles the results and returns an
// error if there is one.
// If too many arguments are provided, the superfluous will be ignored (e.g. if
// a command accepts 1 argument and 2 are provided, the second will be
// ignored). If not enough arguments are provided, it returns an error.
func HandleCmd(s *session.Session, c *gateway.MessageCreateEvent, cmd string, args []string) error {

	// Number of arguments for each command
	nargs := map[string]int {
		"getrpost": 1,
	};

	if len(args) < nargs[cmd] {
		return fmt.Errorf("Not enough arguments for command `" + cmd + "'")
	}

	// Route the commands and handle their result
	switch (cmd) {
	case "getrpost":
		url, err := getRedditPost(args[0])
		if err != nil {
			SendMessage(s, c.ChannelID, err.Error())
		} else {
			SendMessage(s, c.ChannelID, url)
		}
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
	posts, _, err := client.Subreddit.RisingPosts(context.Background(), sub, nil)
	if err != nil {
		return "", err
	}

	postIdx := rand.Intn(len(posts))

	if len(posts) > 0 {
		return "https://reddit.com" + posts[postIdx].Permalink, nil
	} else {
		return "", fmt.Errorf("No new posts")
	}
}

