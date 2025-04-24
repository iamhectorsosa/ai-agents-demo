package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/eduardolat/openroutergo"
	"github.com/iamhectorsosa/ai-agents-demo/internal/config"
	"github.com/iamhectorsosa/ai-agents-demo/internal/logger"
	"github.com/iamhectorsosa/ai-agents-demo/internal/repository/tools"
)

func runAgent(log *logger.Logger, cfg *config.Config, client *openroutergo.Client, systemPrompt, userPrompt string) {
	// INFO: ⚙️ setup client
	compl := client.NewChatCompletion().
		WithContext(context.Background()).
		WithModel(cfg.Model).
		WithModelFallback(cfg.ModelFallback).
		WithTool(tools.ThinkTool).
		WithTool(tools.PrintEntitiesTool).
		WithTool(tools.AnalyzeSentimentTool).
		WithSystemMessage(systemPrompt)

	// INFO: 👤 capture the user prompt
	message := openroutergo.ChatCompletionMessage{
		Role:    openroutergo.RoleUser,
		Content: userPrompt,
	}
	log.User(message.Content)

	// INFO: 🏁 fire up the agent
	startTime := time.Now()
	cycleCount := 1

	for {
		_, resp, err := compl.
			WithMessage(message).
			Execute()
		if err != nil || !resp.HasChoices() {
			log.Error("Error or empty response", "duration", time.Since(startTime), "err", err)
			break
		}

		log.System("Completion executed", "cycleCount", cycleCount, "duration", time.Since(startTime), "usage", resp.Usage)

		// INFO: 🧐 process the response - check messages
		for _, choice := range resp.Choices {
			if choice.Message.Content == "" {
				continue
			}
			switch role := choice.Message.Role; role {
			case openroutergo.RoleAssistant:
				log.Agent(choice.Message.Content)
			default:
				log.System("message content", "from", role, "content", choice.Message.Content)
			}
		}

		// INFO: 📝 prepare for next LLM execution
		shouldContinue := false
		var inputs any
		draftMessage := openroutergo.ChatCompletionMessage{
			Role: openroutergo.RoleTool,
		}

		// INFO: 🧐 process the response - execute handlers
		for _, choice := range resp.Choices {
			if choice.Message.HasToolCalls() {
				tc := choice.Message.ToolCalls
				for _, tool := range tc {
					switch toolName := tool.Function.Name; toolName {
					case tools.ThinkTool.Name:
						toolArguments := tool.Function.Arguments
						var thought tools.ThoughtInput
						if err := json.Unmarshal([]byte(toolArguments), &thought); err != nil {
							log.Error("Unmarshal thought input", "duration", time.Since(startTime), "err", err)
							break
						}
						shouldContinue, inputs = true, thought
						draftMessage.ToolCallID, draftMessage.Name, draftMessage.Content = tool.ID, toolName, toolArguments

					case tools.PrintEntitiesTool.Name:
						toolArguments := tool.Function.Arguments
						var entities tools.EntitiesInput
						if err := json.Unmarshal([]byte(toolArguments), &entities); err != nil {
							log.Error("Unmarshal entities input", "duration", time.Since(startTime), "err", err)
							break
						}
						shouldContinue, inputs = true, entities
						draftMessage.ToolCallID, draftMessage.Name, draftMessage.Content = tool.ID, toolName, toolArguments

					case tools.AnalyzeSentimentTool.Name:
						toolArguments := tool.Function.Arguments
						var sentiment tools.SentimentAnalysisInput
						if err := json.Unmarshal([]byte(toolArguments), &sentiment); err != nil {
							log.Error("Unmarshal sentiment input", "duration", time.Since(startTime), "err", err)
							break
						}
						shouldContinue, inputs = true, sentiment
						draftMessage.ToolCallID, draftMessage.Name, draftMessage.Content = tool.ID, toolName, toolArguments

					default:
						log.Warning("unexpected tool call", "toolName", toolName, "inputs", tool.Function.Arguments)
					}
				}
			}
		}

		// INFO: 🧐 evaluate next execution
		if !shouldContinue {
			log.System("No further requests, exiting...", "duration", time.Since(startTime))
			break
		}

		// INFO: 📢 report handlers
		log.System("Tool executed", "duration", time.Since(startTime), "toolName", draftMessage.Name, "inputs", inputs)
		message = draftMessage
		cycleCount++
	}
}
