package main

import (
	"context"
	"os"
	"log"
	"strings"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

// Read token from filesystem
func readToken() (string, error) {
	token, err := os.ReadFile("DISCORD_TOKEN")

	if err == nil {
		// Convert to string and remove spaces (such as EOF)
		return strings.TrimSpace(string(token)), nil
	}

	return "", err
}

func main() {
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
		log.Println(c.Author.Username, "sent", c.Content)
	})

	// Block forever
	select {}
}

