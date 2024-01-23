package main

import (
	"fmt"
	"os"

	"github.com/lanseg/tgbot"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <bot_api_key>")
		os.Exit(1)
	}
	apiKey := os.Args[1]
	api := tgbot.NewTelegramApi(tgbot.NewBot(apiKey))
	result, err := api.GetUpdates(&tgbot.GetUpdatesRequest{
		Timeout: 60, // 60 Seconds
	})
	if err != nil {
		fmt.Printf("Could not perform request: %s\n", err)
		os.Exit(1)
	}

	for _, upd := range result.Result {
		msg := upd.Message
		fmt.Printf("Got new message %d in chat %d: %s\n", msg.MessageID, msg.Chat.ID, msg.Text)
		result, err := api.SetMessageReaction(&tgbot.SetMessageReactionRequest{
			MessageID: msg.MessageID,
			ChatID:    fmt.Sprintf("%d", msg.Chat.ID),
			Reaction: []*tgbot.ReactionType{
				{
					Type:  "emoji",
					Emoji: "❤",
				},
			},
		})
		fmt.Printf("Got a response: %s: %s", result, err)
	}
}
