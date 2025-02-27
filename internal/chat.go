package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mnafiudrr/techat/internal/ollama"
	"github.com/mnafiudrr/techat/internal/telegram"
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
		Message struct {
			Chat struct {
				ID int64 `json:"id"`
			} `json:"chat"`
			Text string `json:"text"`
		} `json:"message"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	select {
	case workerPool <- struct{}{}:
		go processTelegramMessage(request.Message.Chat.ID, request.Message.Text)
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

func processTelegramMessage(chatID int64, text string) {
	defer func() { <-workerPool }()

	command := telegram.GetCommandArguments(text)
	if command == "start" {
		telegram.CommandStart(chatID)
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
