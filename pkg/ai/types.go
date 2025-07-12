package ai

// InstagramContent represents the LLM output for Instagram posts
// Slides: short texts for each slide
// Caption: detailed, SEO-optimized caption
type InstagramContent struct {
	Slides  []string `json:"slides"`
	Caption string   `json:"caption"`
}
