package tgbot

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <bot_api_key>")
		os.Exit(1)
	}
	apiKey := os.Args[1]
	bot := NewBot(apiKey)
	result, err := bot.GetUpdates(&GetUpdatesRequest{})
	if err != nil {
		fmt.Printf("Could not perform request: %s\n", err)
		os.Exit(1)
	}

	asJson, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Got an update: %s\n", asJson)
}
