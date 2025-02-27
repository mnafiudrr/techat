package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mnafiudrr/techat/internal"
)

func routes(r *gin.Engine) {
	apiRequest := r.Group("/api")

	apiRequest.POST("/chat", internal.Chat)
	apiRequest.POST("/telegram/webhook", internal.TelegramWebhook)
}
