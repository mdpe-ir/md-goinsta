package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mdpe-ir/md-goinsta/internal/config"
)

// LLMClient is a simple HTTP client for OpenRouter

type LLMClient struct {
	APIKey  string
	Referer string
	Title   string
}

func NewLLMClient() *LLMClient {
	cfg := config.Load()
	return &LLMClient{
		APIKey:  cfg.Ai.ApiKey,
		Referer: "", // customize as needed
		Title:   "", // customize as needed
	}
}

// SendPrompt sends a prompt to the OpenRouter API and returns the raw response string
func (c *LLMClient) SendPrompt(model, prompt string) (string, error) {
	requestBody := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{{
			"role":    "user",
			"content": prompt,
		}},
	}
	bodyBytes, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	if c.Referer != "" {
		req.Header.Set("HTTP-Referer", c.Referer)
	}
	if c.Title != "" {
		req.Header.Set("X-Title", c.Title)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var apiResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(respBytes, &apiResp); err != nil {
		return "", err
	}
	if len(apiResp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from LLM")
	}
	return apiResp.Choices[0].Message.Content, nil
}
