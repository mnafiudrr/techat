package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const telegramAPI = "https://api.telegram.org/bot%s/sendMessage"

type sendMessageRequestBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessage(ChatID int64, Text string) error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable is not set")
	}

	requestBody := &sendMessageRequestBody{
		ChatID: ChatID,
		Text:   Text,
	}

	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf(telegramAPI, botToken),
		"application/json",
		bytes.NewBuffer(requestBytes),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	fmt.Println("Message sent to chat ID:", ChatID)

	return nil
}
