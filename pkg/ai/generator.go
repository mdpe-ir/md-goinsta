package ai

import (
	"encoding/json"
	"fmt"
)

// Remove triple backticks and optional 'json' label
func stripMarkdownJSON(resp string) string {
	// Remove triple backticks and optional 'json' label
	if len(resp) > 0 {
		if idx := len(resp); idx > 0 {
			// No operation needed here
		}
		if len(resp) > 0 && resp[:3] == "```" {
			resp = resp[3:]
			if len(resp) > 4 && resp[:4] == "json" {
				resp = resp[4:]
			}
			// Remove leading newlines/spaces
			for len(resp) > 0 && (resp[0] == '\n' || resp[0] == ' ' || resp[0] == '\r') {
				resp = resp[1:]
			}
			// Remove trailing backticks
			if i := len(resp) - 3; i > 0 && resp[i:] == "```" {
				resp = resp[:i]
			}
		}
	}
	return resp
}

// GenerateInstagramContent uses LLMClient to generate InstagramContent from a subject
func GenerateInstagramContent(client *LLMClient, subject, lang string) (*InstagramContent, error) {
	prompt := fmt.Sprintf(`You are an expert social media content creator and Instagram strategist. Your task is to generate engaging, educational content for an Instagram carousel post based on the provided subject.

Instructions:
- Return a JSON object with two fields:
  - "slides": an array of 5-7 concise, visually appealing texts for Instagram carousel slides. Each slide must:
    - Be 8-12 words for brevity and readability.
    - Use catchy, clear language suitable for overlaying on images.
    - Include hooks, facts, or questions to encourage swiping.
  - "caption": an SEO-optimized Instagram caption that:
    - Expands on the slide content with engaging, informative details.
    - Uses 100-200 words, formatted in short paragraphs with emojis and line breaks for readability.
    - Includes 5-10 relevant hashtags to boost discoverability.
    - Ends with a strong call-to-action (CTA) to drive engagement (e.g., likes, comments, shares, or follows).
    - Aligns with Instagram’s latest algorithm and SEO best practices (e.g., keyword-rich, engaging tone).
- Write all output in %s language.
- Ensure content is accurate, relevant to the subject, and optimized for Instagram’s visual and social format.

Subject: %s

Sample output:
{
  "slides": [
    "What is Golang?",
    "Created by Google in 2009",
    "Fast, efficient, statically typed",
    "Ideal for cloud and backend",
    "Thriving developer community",
    "Concurrency made easy with goroutines",
    "Ready to learn Go?"
  ],
  "caption": "🚀 Discover Golang (Go), a powerful language built by Google! Known for speed, simplicity, and reliability, Go is perfect for cloud services, APIs, and scalable systems. Its clean syntax and goroutines make concurrency a breeze, winning over backend developers worldwide. 🌐\n\nWhether you're a beginner or a pro, Go’s growing community and demand in tech make it a must-learn skill. Start coding today! 💻\n\nWhat’s your favorite programming language? Share below! ⬇️\n\n#golang #programming #backenddevelopment #cloudcomputing #learntocode #developers #codinglife #techskills"
}

Output only the JSON object, nothing else.`, lang, subject)

	resp, err := client.SendPrompt("deepseek/deepseek-chat-v3-0324:free", prompt)
	if err != nil {
		return nil, err
	}
	fmt.Println("LLM raw response:\n", resp)

	resp = stripMarkdownJSON(resp)

	var result InstagramContent
	if err := json.Unmarshal([]byte(resp), &result); err != nil {
		return nil, fmt.Errorf("failed to parse InstagramContent: %v", err)
	}
	return &result, nil
}
