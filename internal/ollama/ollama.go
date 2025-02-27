package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OllamaRequest struct {
	Model          string  `json:"model"`
	Prompt         string  `json:"prompt"`
	NumPredictions int     `json:"num_predictions"`
	MaxTokens      int     `json:"max_tokens"`
	Temperature    float64 `json:"temperature"`
	TopP           float64 `json:"top_p"`
	Stream         bool    `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func RequestPrompt(prompt string) (string, error) {
	ollamaURL := os.Getenv("OLLAMA_URL")
	if ollamaURL == "" {
		return "", fmt.Errorf("OLLAMA_URL environment variable is not set")
	}

	model := os.Getenv("OLLAMA_MODEL")
	if model == "" {
		model = "qwen:1.8b"
	}

	templatePrompt := os.Getenv("OLLAMA_TEMPLATE_PROMPT")
	if templatePrompt != "" {
		prompt = fmt.Sprintf(templatePrompt, prompt)
	}

	requestBody, err := json.Marshal(OllamaRequest{
		Model:          model,
		Prompt:         prompt,
		NumPredictions: 1,
		MaxTokens:      100,
		Temperature:    0.5,
		TopP:           0.9,
		Stream:         false,
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Response, nil
}
