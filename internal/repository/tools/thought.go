package tools

import "github.com/eduardolat/openroutergo"

type ThoughtInput struct {
	Thought string `json:"thought"`
}

var ThinkTool = openroutergo.ChatCompletionTool{
	Name:        "think_tool",
	Description: "Use the tool to think about something. It will not obtain new information, but use it only when complex reasoning is needed and you feel that you need to take a moment to think about the user's request.",
	Parameters: map[string]any{
		"type": "object",
		"properties": map[string]any{
			"thought": map[string]any{
				"type":        "string",
				"description": "A thought to think about.",
			},
		},
		"required": []string{"thought"},
	},
}
