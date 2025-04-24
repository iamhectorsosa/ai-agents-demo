package tools

import "github.com/eduardolat/openroutergo"

type EntitiesInput struct {
	Entities []struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Context string `json:"context"`
	} `json:"entities"`
}

var PrintEntitiesTool = openroutergo.ChatCompletionTool{
	Name:        "print_entities_tool",
	Description: "Extract all named entities provided only by the user before generating a response.",
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
