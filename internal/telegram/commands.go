package telegram

import "github.com/mnafiudrr/techat/internal/ollama"

func GetCommandArguments(text string) string {
	if len(text) == 0 {
		return ""
	}

	if text[0] == '/' {
		return text[1:]
	}

	return ""
}

func CommandStart(chatID int64) {
	greeting, err := ollama.RequestPrompt("make a greeting that you're a telegram bot that can asked about whatever user want but currenty only accept word and under development. the greeting length is maximum in 10 words. the greeting should be friendly and welcoming, example: 'Hello! I'm a bot. I can help you with stuff.'")
	if err != nil {
		SendMessage(chatID, "Hello! I'm a bot. I can help you with stuff.")
		return
	}
	SendMessage(chatID, greeting+"\n\nHave any idea? Lemme know at @fuifiu")
}
