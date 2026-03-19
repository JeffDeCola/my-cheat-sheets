# OPENCLAW AUTONOMOUS AI-AGENT CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_OpenClaw is a personal AI assistant that runs on your server,
connects to your apps, and does things for you automatically._

tl;dr

```text

```

Table of contents

* tbd

## OVERVIEW

The Stack

```text
      OpenClaw                  The orchestrator - listens, acts, manages (not AI)
      Ollama                    The inference server - hosts and serves the AI model
      llama, qwen, deepseek     The AI model — where the actual intelligence lives
      P40                       The hardware accelerator - for the AI model
```

How OpenClaw Fits in — Request Flow

```text
You ask a question
        ↓
OpenClaw sends it to Ollama via API (POST 192.168.20.141:11434)
        ↓
Ollama loads the requested AI model (llama3.1:8b, etc...) into P40's 24GB VRAM
        ↓
AI Model thinks and generates a response
        ↓
Ollama sends response back to OpenClaw
        ↓
OpenClaw takes action based on the answer
```

> Note: Ollama keeps the model warm in VRAM for a few minutes after
> each request — so the second question is always faster than the first.

## INSTALL

I installed OpenClaw on my dell rack server that has a nvidia Tesla P40 in it.


