package main

import (
	"flag"

	"github.com/eduardolat/openroutergo"
	"github.com/iamhectorsosa/ai-agents-demo/internal/config"
	"github.com/iamhectorsosa/ai-agents-demo/internal/logger"
)

func main() {
	log := logger.New()
	cfg, err := config.New()
	if err != nil {
		log.Error("loading config", "err", err)
		return
	}

	client, err := openroutergo.
		NewClient().
		WithAPIKey(cfg.OpenRouterAPIKey).
		Create()
	if err != nil {
		log.Error("initiating client", "err", err)
		return
	}

	systemPrompt := `
  Always reply in plain text format, always in a single line. 
  Please help the user follow along with details of your thought process.
  Help the user understand your actions and the reason why you called the tools you used if any.
  `
	userPrompt := "Write me a short and concise rhyme about Jozef and Eric from Webscope. They met last week in Prague."

	var demo string
	flag.StringVar(&demo, "demo", "", "")
	flag.Parse()

	switch demo {
	case "workflow":
		runWorkflow(log, cfg, client, systemPrompt, userPrompt)
	case "agent":
		runAgent(log, cfg, client, systemPrompt, userPrompt)
	default:
		log.System("Please select a valid demo...")
	}
}
