package main

import (
	"context"
	"os"
	"log"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

// Read token from filesystem
func readToken() (string, error) {
	tokenBytes, err := os.ReadFile("DISCORD_TOKEN")

	if err == nil {
		token := string(tokenBytes)

		// Remove last character (EOF) from token
		return token[:len(token)-1], nil
	}

	return "", err
}

func main() {
	// The prefix is a string of one or more characters that is used to "call"
	// the bot, which means that the bot will interpret all the strings that
	// start with that prefix as commands.
	// Discord will replace the needs of prefixes with a common prefix "/" so
	// this is just a temporary solution.
	prefix := "!"

	token, err := readToken()
	if err != nil || token == "" {
		log.Fatalln("Cannot read file `DISCORD_TOKEN' or token is empty")
	}

	log.Println("File `DISCORD_TOKEN' found")

	// Create session with intents
	s, err := session.NewWithIntents(token, gateway.IntentGuildMessages,
		gateway.IntentGuilds)
	if err != nil {
		log.Fatalln("Cannot create a session, check the error below.\n", err)
	}

	if err := s.Open(context.Background()); err != nil {
		log.Fatalln("Failed to connect, check the error below.\n", err)
	}
	defer s.Close()

	log.Println("Session created")

	s.AddHandler(func(c *gateway.MessageCreateEvent) {
		command := strings.TrimPrefix(c.Content)

		if command == c.Content {
			// It the content remains unchanged, there was no prefix
			return
		}


		//log.Println(c.Author.Username, "sent", c.Content)
	})

	log.Println("Bot has successfully started, use CTRL-C to stop it")
	// Block forever
	select {}
}

