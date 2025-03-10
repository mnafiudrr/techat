package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mnafiudrr/techat/internal/ollama"
	"github.com/mnafiudrr/techat/internal/telegram"
	"github.com/mnafiudrr/techat/internal/telegram/types"
)

var workerPool = make(chan struct{}, 50)

func Chat(c *gin.Context) {
	var request struct {
		Prompt string `json:"prompt"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	select {
	case workerPool <- struct{}{}:
		go processChat(request.Prompt)
		c.JSON(http.StatusOK, gin.H{"message": "Processing your request asynchronously!", "prompt": request.Prompt})
	default:
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Server busy, please try again later."})
	}
}

func TelegramWebhook(c *gin.Context) {
	var request struct {
		types.WebhookRequestType
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	select {
	case workerPool <- struct{}{}:
		go processTelegramMessage(request.Message.Chat.ID, *request.Message.Text, request.WebhookRequestType)
		c.JSON(http.StatusOK, gin.H{"message": "Processing your request asynchronously!", "chat_id": request.Message.Chat.ID, "text": request.Message.Text})
	default:
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Server busy, please try again later."})
	}
}

func processChat(prompt string) {
	defer func() { <-workerPool }()

	fmt.Println("Processing:", prompt)
	response, err := ollama.RequestPrompt(prompt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", response)
}

func processTelegramMessage(chatID int64, text string, request types.WebhookRequestType) {
	sendBackTheRequest := os.Getenv("SEND_BACK_THE_REQUEST")
	if sendBackTheRequest == "true" {
		requestString, _ := json.Marshal(request)
		text = fmt.Sprintf("```json\n%s\n```", requestString)
		telegram.SendMessage(chatID, text)
	}

	defer func() { <-workerPool }()

	command := telegram.GetCommandArguments(text)
	if command == "start" {
		telegram.CommandStart(chatID)
		return
	}

	if command == "ping" {
		telegram.SendMessage(chatID, "Pong!")
		return
	}

	if command != "" {
		telegram.SendMessage(chatID, "I don't understand that command.")
		return
	}

	fmt.Println("Processing:", text)
	response, err := ollama.RequestPrompt(text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", response)
	telegram.SendMessage(chatID, response)
}
