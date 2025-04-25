package main

import (
	"context"
	"time"

	"github.com/eduardolat/openroutergo"
	"github.com/iamhectorsosa/ai-agents-demo/internal/config"
	"github.com/iamhectorsosa/ai-agents-demo/internal/logger"
)

func runWorkflow(
	log *logger.Logger,
	cfg *config.Config,
	client *openroutergo.Client,
	systemPrompt,
	userPrompt string,
) {
	// INFO: ⚙️ setup client
	compl := client.NewChatCompletion().
		WithContext(context.Background()).
		WithModel(cfg.Model).
		WithModelFallback(cfg.ModelFallback).
		WithSystemMessage(systemPrompt)

	// INFO: 👤 capture the user prompt
	message := openroutergo.ChatCompletionMessage{
		Role:    openroutergo.RoleUser,
		Content: userPrompt,
	}
	log.User(message.Content)

	// INFO: 🏁 fire up the workflow
	startTime := time.Now()

	_, resp, err := compl.
		WithMessage(message).
		Execute()
	if err != nil || !resp.HasChoices() {
		log.Error("Error or empty response", "duration", time.Since(startTime), "err", err)
		return
	}

	log.System("Completion executed", "duration", time.Since(startTime), "usage", resp.Usage)

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

	log.System("No further requests, exiting...", "duration", time.Since(startTime))
}
