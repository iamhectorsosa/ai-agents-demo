package main

import "github.com/eduardolat/openroutergo"

type EntitiesInput struct {
	Entities []struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Context string `json:"context"`
	} `json:"entities"`
}

type SentimentAnalysisInput struct {
	Sentiment string  `json:"sentiment"`
	Score     float64 `json:"score"`
}

var PrintEntitiesTool = openroutergo.ChatCompletionTool{
	Name:        "print_entities_tool",
	Description: "Extract all named entities provided only by the user before generating a response",
	Parameters: map[string]any{
		"type": "object",
		"properties": map[string]any{
			"entities": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"name":    map[string]any{"type": "string", "description": "The extracted entity name"},
						"type":    map[string]any{"type": "string", "enum": []string{"PERSON", "ORGANIZATION", "LOCATION"}, "description": "The extracted entity type"},
						"context": map[string]any{"type": "string", "description": "The context in which the entity appears in the text"},
					},
					"required": []string{"name", "type", "context"},
				},
			},
		},
		"required": []string{"entities"},
	},
}

var AnalyzeSentimentTool = openroutergo.ChatCompletionTool{
	Name:        "anaylze_sentiment_tool",
	Description: "Analyze the sentiment of the response generated",
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
