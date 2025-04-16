# AI Agents Demo

![Demo](./demo.gif)

## Overview

This project demonstrates an AI Agent system that can perform agentic processes - autonomous, goal-directed actions using large language models (LLMs). The agent can understand context, make decisions, and use specialized tools to accomplish tasks.

This demo showcases how an LLM can function as an agent by:
- Processing input messages
- Deciding which tools to use
- Executing those tools
- Continuing the conversation with the results

## Project Structure

```
.
├── main.go         # Entry point and main execution logic
├── logger.go       # Logging utilities for system, user, and agent messages
├── config.go       # Configuration management
├── tools.go        # Tool definitions (entity extraction, sentiment analysis, etc.)
```

In a production environment, these components would likely be organized into proper modules:
- `logger/` - Structured logging system
- `tools/` - Tool definitions and implementations
- `config/` - Configuration management
- `agent/` - Agent orchestration logic


## How It Works

The system follows this general flow:
1. User input is captured
2. The LLM processes the input with a system prompt that guides its behavior
3. Tools are executed based on the LLM's decisions:
   - Entity extraction to identify key elements in text
   - Sentiment analysis to evaluate emotional tone
4. Results are formatted and returned to the user
5. The process continues in cycles until completion

The demo uses a structured logging system to clearly display the interaction between the user, system, and AI agent, making it easy to follow the agentic process.
