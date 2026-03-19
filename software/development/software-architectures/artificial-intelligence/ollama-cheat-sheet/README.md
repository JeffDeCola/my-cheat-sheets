# OLLAMA CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Ollama is a local LLM server that lets you download
and run AI models like Llama, Mistral, and others on
your own hardware via a simple API."_

Table of contents

* tbd

## OVERVIEW

How Ollama Works — Request Flow

```text
You ask a question
        ↓
OpenClaw sends it to Ollama via API (POST 192.168.20.141:11434)
        ↓
Ollama loads the requested model into P40's 24GB VRAM
        ↓
Model thinks and generates a response
        ↓
Ollama sends response back to OpenClaw
        ↓
OpenClaw takes action based on the answer
```

> Note: Ollama keeps the model warm in VRAM for a few minutes after
> each request — so the second question is always faster than the first.

## INSTALL

We will install ollama on my dell rack server that has a nvidia Tesla P40 in it.
