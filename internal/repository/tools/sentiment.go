package tools

import "github.com/eduardolat/openroutergo"

type SentimentAnalysisInput struct {
	Sentiment string  `json:"sentiment"`
	Score     float64 `json:"score"`
}

var AnalyzeSentimentTool = openroutergo.ChatCompletionTool{
	Name:        "anaylze_sentiment_tool",
	Description: "Analyze the sentiment of the response generated. This tool should only be ran at the end of generated text to provide insight on the Agent/LLM response.",
	Parameters: map[string]any{
		"type": "object",
		"properties": map[string]any{
			"sentiment": map[string]any{
				"type":        "string",
				"enum":        []string{"positive", "neutral", "negative"},
				"description": "The overall sentiment of the response",
			},
			"score": map[string]any{
				"type":        "number",
				"minimum":     0,
				"maximum":     1,
				"description": "Confidence score (0-1) for the sentiment analysis",
			},
		},
		"required": []string{"sentiment", "score"},
	},
}
