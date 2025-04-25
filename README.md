# AI Agents Demo

![Demo](./demo.gif)

## Requirements

* Go 1.24.0
* OpenRouter API Key

For environment variables please check the example file: [./.env.example](./.env.example)

## Overview

This project demonstrates an AI Agent system that leverages large language models (LLMs) through the OpenRouter API to perform autonomous, goal-directed actions. The agent can understand context, make decisions, and use specialized tools to accomplish tasks.

The application supports two execution modes:
- **Workflow Mode**: A simple, single-cycle interaction with the LLM
- **Agent Mode**: A multi-cycle interaction where the LLM can use tools and continue the conversation

## Project Structure

```
.
├── main.go                         # Entry point and execution logic
├── run_workflow.go                 # Simple workflow implementation
├── run_agent.go                    # Interactive agent implementation
├── internal/
│   ├── config/                     # Configuration management
│   ├── logger/                     # Structured logging system
│   └── repository/
│       └── tools/                  # Tool definitions and implementations
```

## How It Works

The system follows this general flow:
1. User input is captured
2. The LLM processes the input with a system prompt that guides its behavior
3. In agent mode, tools are executed based on the LLM's decisions:
   - Entity extraction to identify key elements in text
   - Sentiment analysis to evaluate emotional tone
   - Thinking tool to reason through complex problems
4. Results are formatted and returned to the user
5. In agent mode, the process continues in cycles until completion

## Running the Demo

To run the workflow demo:
```
go run . --demo workflow
```

To run the agent demo:
```
go run . --demo agent
```

The demo uses a structured logging system to clearly display the interaction between the user, system, and AI agent, making it easy to follow the agentic process.
