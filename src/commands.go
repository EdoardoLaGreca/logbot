package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/diamondburned/arikawa/v3/session"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

// A struct for commands
// It contains the name of a command, its description and the number of
// arguments that it requires to be executed
type Command struct {
	name	string
	desc	string
	usage	string
	reqArgs	uint
}

// Available commands
var commands = map[string]Command {
	"help": Command{ "help", "Show a help page", "help", 0 },
	"grp": Command{ "grp", "Get Reddit post", "grp <subreddit>", 1 },
}


// Routes the commands to their routines, handles the results and returns an
// error if there is one.
// If too many arguments are provided, the superfluous will be ignored (e.g. if
// a command accepts 1 argument and 2 are provided, the second will be
// ignored). If not enough arguments are provided, it returns an error.
func HandleCmd(s *session.Session, c *gateway.MessageCreateEvent, cmd string, args []string) error {

	if len(args) < int(commands[cmd].reqArgs) {
		return fmt.Errorf("Not enough arguments for command `" + cmd + "'")
	}

	// Route the commands and handle their result
	switch (cmd) {
	case "help":
		helpPage := help()
		sendMessage(s, c.ChannelID, helpPage)
	case "grp":
		url, err := getRedditPost(args[0])
		if err != nil {
			sendMessage(s, c.ChannelID, err.Error())
		} else {
			sendMessage(s, c.ChannelID, url)
		}
	default:
		return fmt.Errorf("Unknown command `" + cmd + "'")
	}

	return nil
}

func help() string {
	helpPage := "```"

	// Header
	helpPage += "HELP PAGE\n"

	// Size in characters of the longest command name, used for formatting
	longestCmd := 0

	// Get the longest command name
	for _, cmd := range commands {
		if len(cmd.name) > longestCmd {
			longestCmd = len(cmd.name)
		}
	}

	// Add command info to helpPage for each command
	for _, cmd := range commands {
		helpPage += " " + cmd.name

		// Add spacing
		lenDiff := longestCmd - len(cmd.name) + 2
		for i := 0; i < lenDiff; i++ {
			helpPage += " "
		}

		helpPage += cmd.desc + "\n   Usage: " + cmd.usage + "\n\n"
	}

	helpPage += "```"

	return helpPage
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

