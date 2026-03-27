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
ollama run llama3.1:8b --verbose                     # shows info like tokens
ollama run llama3.1:8b "explain neural networks"     # one-shot prompt
ollama ps                                            # models loaded in VRAM
ollama show llama3.1:8b                              # model info
ollama stop llama3.1:8b                              # unload model from VRAM
curl http://localhost:11434                          # check Ollama is running
```

Table of contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#overview)
  * [The Stack](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#the-stack)
  * [Claude vs Homelab](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#claude-vs-homelab)
  * [How Ollama Fits In — Request Flow](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#how-ollama-fits-in--request-flow)
  * [What OpenClaw Actually Does](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#what-openclaw-actually-does)
  * [Picking Your LLM](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#picking-your-llm)
  * [Feeding Knowledge Into Your LLM](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#feeding-knowledge-into-your-llm)
* [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#install)
* [MODELS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#models)
  * [Models I Have](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#models-i-have)
  * [What Model to Use](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#what-model-to-use)
  * [Tokens/s on P40](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#tokenss-on-p40)
* [CONFIGURE OLLAMA TO LISTEN ON ALL INTERFACES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-inference-stack/ollama-cheat-sheet#configure-ollama-to-listen-on-all-interfaces)

## OVERVIEW

### The Stack

```text
      OpenClaw / Open WebUI     The orchestrator - listens, acts, manages (not AI)
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

### Claude vs Homelab

```text
                    CLAUDE                  HOMELAB

LLM                 Claude (Anthropic)      LLaMA, DeepSeek, Qwen, etc...
Inference server    Anthropic cloud         Ollama
Hosting             Anthropic cloud         Your P40
Agent               Built into Claude       OpenClaw
Interface           claude.ai               Telegram, iPhone app, etc.
Memory              Per session only        Persistent (OpenClaw)
Privacy             Data goes to Anthropic  100% stays on your hardware
Cost                Subscription            Free after hardware investment
Internet access     Yes (web search)        Via OpenClaw tools
Knowledge cutoff    Aug 2025                Depends on model
Control             None                    100% yours
```

> This whole thing is called a Self-hosted AI Stack —
> running AI on your own hardware instead of the cloud.

### How Ollama Fits In — Request Flow

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

### What OpenClaw Actually Does

An AI agent does more than feed knowledge — it:

* **Feeds knowledge**    — pulls real-time context (weather, files, web)
* **Takes actions**      — sends emails, creates events, runs code
* **Remembers**          — persistent memory across sessions
* **Plans**              — breaks tasks into steps and executes them
* **Runs autonomously**  — wakes up on a schedule without you asking

### Picking Your LLM

You pick your LLM (frozen brain) based on a few things:

* **Different training data**      — some focused on code, some on science,
                                     some on general knowledge
* **Different sizes**              — more parameters = more capacity
                                     to store knowledge and reason
* **Different specializations**    — some are tuned for reasoning,
                                     some for speed, some for tool calling

### Feeding Knowledge Into Your LLM

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

## MODELS

Pull some models

```bash
ollama pull qwen3:8b             # quick tests + lightweight fallback (~5GB)
```

verify

```bash
ollama list
```

### Models I Have

**Updated March 2026**

| Model | Parameters | Disk Size | VRAM (Q4) | Context Window | Reasoning | Tool Calling | Use Case | Notes |
|-------|-----------|-----------|-----------|----------------|-----------|-------------|----------|-------|
| **phi4-mini** | 3.8B | 2.5 GB | ~3 GB | 16K | No | No | Fast chat, quick Q&A | Fastest model, best for simple tasks |
| **llama3.1:8b** | 8B | 4.9 GB | ~5 GB | 128K | No | Yes | General purpose, everyday tasks | Meta model, huge community, solid all-rounder |
| **qwen3:8b** | 8B | 5.2 GB | ~5 GB | 40K | Yes (/think) | Yes | Agents, tool calling, fallback | Default for OpenClaw agents on VM 103/104 |
| **qwen2.5-coder:32b** | 32B | 19 GB | ~20 GB | 32K | No | Yes | Coding, refactoring, code review | Best local coding model, fills P40 VRAM |
| **glm-4.7-flash** | ~9B | 19 GB | ~12 GB | 1M | No | No | Long context analysis | Huge context window, good for large documents |
| **deepseek-r1:32b** | 32B | 19 GB | ~20 GB | 128K | Yes (CoT) | **No** | Deep reasoning, math, complex analysis | Chain-of-thought built in, NO tool calling |

### What Model to Use

* **Fastest response:** phi4-mini
* **Best all-rounder:** llama3.1:8b
* **Best for agents:** qwen3:8b (tool calling + reasoning)
* **Best for coding:** qwen2.5-coder:32b
* **Best for long docs:** glm-4.7-flash (1M context)
* **Best for hard problems:** deepseek-r1:32b (but no tool calling)

### Tokens/s on P40

| Model | Prompt Eval | Generation |
|-------|-------------|------------|
| phi4-mini | ~800+ t/s | ~50-65 t/s |
| llama3.1:8b | ~500+ t/s | ~30-45 t/s |
| qwen3:8b | ~450+ t/s | ~25-42 t/s |
| qwen2.5-coder:32b | ~150+ t/s | ~12-18 t/s |
| glm-4.7-flash | ~300+ t/s | ~20-30 t/s |
| deepseek-r1:32b | ~150+ t/s | ~12-18 t/s |

## CONFIGURE OLLAMA TO LISTEN ON ALL INTERFACES

By default Ollama only listens on localhost (127.0.0.1).
To allow other VMs on your network to connect you need to
configure it to listen on all interfaces (0.0.0.0).

### Check What Ollama is Currently Listening On

```bash
ss -tlnp | grep 11434
```

Expected output after fix

```bash
LISTEN 0      4096               *:11434            *:*
```

Set OLLAMA_HOST environment variable

```bash
sudo systemctl edit ollama
```

Add these lines

```text
[Service]
Environment="OLLAMA_HOST=0.0.0.0:11434"
```

```bash
sudo systemctl daemon-reload
sudo systemctl restart ollama
ss -tlnp | grep 11434
```

Verify from another VM

```bash
curl http://192.168.20.141:11434
# Should return: Ollama is running
```
