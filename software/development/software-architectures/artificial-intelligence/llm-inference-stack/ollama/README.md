# OLLAMA CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Ollama is a local LLM server that lets you download
and run AI models (LLMs) like Llama, Mistral, and others on
your own hardware via a simple API."_

tl;dr

```text
ollama list                                          # list all downloaded models
ollama pull llama3.1:8b                              # download a model
ollama rm llama3.1:8b                                # delete a model
ollama run llama3.1:8b                               # interactive chat
ollama run llama3.1:8b "explain neural networks"     # one-shot prompt
ollama ps                                            # models loaded in VRAM
ollama show llama3.1:8b                              # model info
ollama stop llama3.1:8b                              # unload model from VRAM
curl http://localhost:11434                          # check Ollama is running
```

Table of contents

* tbd

## OVERVIEW

### The Stack

```text
      OpenClaw                  The orchestrator - listens, acts, manages (not AI)
      Ollama                    The inference server - hosts and serves the AI model
      LLM                       The AI model — where the actual intelligence lives
        llama3.1:8b             Meta        8B params   ~5GB   Fast, lightweight
        deepseek-r1:32b         DeepSeek   32B params  ~19GB   Heavy reasoning
        qwen2.5-coder:32b       Alibaba    32B params  ~19GB   Coding + tool calling
        qwen3:8b                Alibaba     8B params   ~5GB   Quick tests + fallback
      Nvidia P40                The hardware accelerator - powers LLMs - 24GB VRAM
```

> LLMs are frozen brains — the knowledge is baked in during training.
> They do not connect to the internet or seek new information.
> Their knowledge stops at their training cutoff date.

### CLAUDE VS HOMELAB

```text
                    Claude              Your Homelab

LLM                 Claude (Anthropic)  LLaMA, DeepSeek, Qwen
Inference server    Anthropic cloud     Ollama
Hosting             Anthropic cloud     Your P40
Agent               Built into Claude   OpenClaw
Interface           claude.ai           Telegram, iPhone app, etc.
Memory              Per session only    Persistent (OpenClaw)
Privacy             Data goes to        100% stays on your
                    Anthropic           hardware
Cost                Subscription        Free after hardware
Internet access     Yes (web search)    Via OpenClaw tools
Knowledge cutoff    Aug 2025            Depends on model
Control             None                100% yours
```

> This whole thing is called a Self-hosted AI Stack —
> running AI on your own hardware instead of the cloud.

### HOW OLLAMA FITS IN — REQUEST FLOW

```text
You ask a question
        ↓
OpenClaw sends it to Ollama via API (POST 192.168.20.141:11434)
        ↓
Ollama loads the requested AI model (llama3.1:8b, etc...) into P40's 24GB VRAM
        ↓
AI Model (LLM) thinks and generates a response
        ↓
Ollama sends response back to OpenClaw
        ↓
OpenClaw takes action based on the answer
```

> Note: Ollama keeps the model warm in VRAM for a few minutes after
> each request — so the second question is always faster than the first.

### WHAT OPENCLAW ACTUALLY DOES

An AI agent does more than feed knowledge — it:

* **Feeds knowledge**    — pulls real-time context (weather, files, web)
* **Takes actions**      — sends emails, creates events, runs code
* **Remembers**          — persistent memory across sessions
* **Plans**              — breaks tasks into steps and executes them
* **Runs autonomously**  — wakes up on a schedule without you asking

### PICKING YOUR LLM

You pick your LLM (frozen brain) based on a few things:

* **Different training data**      — some focused on code, some on science, some on general knowledge
* **Different sizes**              — more parameters = more capacity to store knowledge and reason
* **Different specializations**    — some are tuned for reasoning, some for speed, some for tool calling

### FEEDING KNOWLEDGE INTO YOUR LLM

You can feed an LLM knowledge directly. OpenClaw does this automatically.
This is called RAG — Retrieval Augmented Generation.
OpenClaw fetches fresh data and feeds it to the LLM as context.

Example — manually feeding context:
```bash
ollama run llama3.1:8b "Today is March 19 2026. The weather in Waltham MA \
is 52F and cloudy. Jeff's meeting is at 3pm. Given this information, should \
Jeff bring a jacket to his meeting and what time should he start getting \
ready if he needs 20 minutes to get ready?"
```

## INSTALL

I installed ollama on my dell rack server that has a nvidia Tesla P40 in it.

```bash
curl -fsSL https://ollama.com/install.sh | sh
```

## PULL MODELS

Pull some models

```bash
ollama pull deepseek-r1:32b      # heavy reasoning + agent tasks (~20GB)
ollama pull qwen2.5-coder:32b    # coding + tool calling for OpenClaw (~20GB)
ollama pull llama3.1:8b          # fast lightweight everyday responses (~5GB)
ollama pull qwen3:8b             # quick tests + lightweight fallback (~5GB)
```

verify

```bash
ollama list
```
