# OPENCLAW AI-AGENT CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_OpenClaw is a personal AI assistant that runs on your server,
connects to your apps, and does things for you automatically._

tl;dr

```text
SETUP
    openclaw --version                                          # Verify install
    openclaw onboard                                            # Interactive setup wizard
    openclaw models                                             # Show default model and config
    openclaw models set ollama/qwen2.5-coder:32b                # Set default model
GATEWAY
    openclaw gateway start                                      # Start gateway
    openclaw health                                             # Health check
    systemctl --user status openclaw-gateway.service            # Service status
USE
    openclaw tui                                                # Terminal interface
    openclaw agent --agent main --message "..." --local         # One-shot agent call
BROWSER (via SSH tunnel)
    ssh -N -L 18789:127.0.0.1:18789 jeff@192.168.20.143         # Then http://localhost:18789
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#overview)
  * [THE STACK](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#the-stack)
  * [REQUEST FLOW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#request-flow)
  * [WHAT OPENCLAW DOES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#what-openclaw-does)
* [INSTALL AND CONFIGURE OPENCLAW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#install-and-configure-openclaw)
  * [INSTALL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#install)
  * [CONFIGURE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#configure)
  * [START THE GATEWAY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#start-the-gateway)
* [ACCESS OPENCLAW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#access-openclaw)
  * [TERMINAL (TUI)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#terminal-tui)
  * [BROWSER VIA SSH TUNNEL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#browser-via-ssh-tunnel)

Documentation and Reference

* AI Fundamentals
  * [artificial intelligence overview](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/ai-fundamentals/artificial-intelligence-overview-cheat-sheet#artificial-intelligence-overview-cheat-sheet)
  * [neural networks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/ai-fundamentals/neural-networks-cheat-sheet#neural-networks-cheat-sheet)
    ([my-neural-networks](https://github.com/JeffDeCola/my-neural-networks?tab=readme-ov-file#my-neural-networks))
  * [math behind training mlp neural networks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/ai-fundamentals/math-behind-training-mlp-neural-networks-cheat-sheet#math-behind-training-mlp-neural-networks-cheat-sheet)
* Application/Orchestration Layer
  * [ai stack configurations - from chatbots to agents](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/ai-stack-configurations-from-chatbots-to-agents-cheat-sheet#ai-stack-configurations---from-chatbots-to-agents-cheat-sheet)
  * **[openclaw ai agent](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/application-orchestration-layer/openclaw-ai-agent-cheat-sheet#openclaw-ai-agent-cheat-sheet)**
    **YOU ARE HERE**
* Inference Layer
  * [ollama](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/inference-layer/ollama-cheat-sheet#ollama-cheat-sheet)
* LLM Layer
  * [llm](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-layer/llm-cheat-sheet#llm-cheat-sheet)
  * [llm training](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/llm-layer/llm-training-cheat-sheet#llm-training-cheat-sheet)

## OVERVIEW

OpenClaw is the orchestration layer between you and your LLM.
It adds tools, memory, and the ability to take action.

### THE STACK

```text
      OpenClaw                  The orchestrator - listens, acts, manages (not AI)
      Ollama                    The inference server - hosts and serves the AI model
      llama, qwen, deepseek     The AI model — the frozen brain
      Nvidia P40                The hardware accelerator - for the frozen brain
```

### REQUEST FLOW

```text
You ask a question
        ↓
OpenClaw sends it to Ollama via API (POST 192.168.20.141:11434)
        ↓
Ollama loads the requested AI model (llama3.1:8b, etc.) into P40's 24GB VRAM
        ↓
AI Model thinks and generates a response
        ↓
Ollama sends response back to OpenClaw
        ↓
OpenClaw takes action based on the answer
```

> Note: Ollama keeps the model warm in VRAM for a few minutes after
> each request — so the second question is always faster than the first.

### WHAT OPENCLAW DOES

```text
Feeds knowledge    — pulls real-time context (weather, files, web)
Takes actions      — sends emails, creates events, runs code
Remembers          — persistent memory across sessions
Plans              — breaks tasks into steps and executes them
Runs autonomously  — wakes up on a schedule without you asking
```

## INSTALL AND CONFIGURE OPENCLAW

Install OpenClaw on Ubuntu via npm, run the onboarding wizard, then start the gateway service.

### INSTALL

I installed OpenClaw on my Dell rack server (running Ubuntu in
a VM) that has an NVIDIA Tesla P40 in it.

Update Ubuntu,

```bash
sudo apt update && sudo apt upgrade -y
```

Install Node.js v22 via NVM,

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
source ~/.bashrc
nvm install 22
nvm use 22
node --version
```

Install OpenClaw via npm,

```bash
npm install -g openclaw
```

Verify,

```bash
openclaw --version
```

Make sure Ollama is listening on all interfaces (so OpenClaw can
reach it from this VM). See the
[ollama](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/inference-layer/ollama-cheat-sheet#configure-ollama-to-listen-on-all-interfaces)
cheat sheet, CONFIGURE OLLAMA TO LISTEN ON ALL INTERFACES section.

Confirm Ollama is reachable from this VM,

```bash
curl http://192.168.20.141:11434
# Should return: Ollama is running
```

### CONFIGURE

Run the interactive setup wizard,

```bash
openclaw onboard
```

Onboarding selections I used,

```text
Security warning       → Yes, continue
Gateway                → Local gateway (this machine)
Workspace              → default (~/.openclaw/workspace)
Model provider         → Ollama
Ollama base URL        → http://192.168.20.141:11434
Ollama mode            → Local
Default model          → ollama/qwen3:8b   (change later if needed)
Gateway port           → 18789 (default)
Gateway bind           → LAN (0.0.0.0)
Gateway auth           → Token (recommended)
Tailscale exposure     → Off
Gateway token          → blank (auto generate)
Chat channels          → No (configure later)
Web search             → Skip for now
Skills                 → No (configure later)
Hooks                  → Skip for now
Systemd service        → Yes
Gateway runtime        → Node (recommended)
Shell completion       → Yes
```

To change the default model later,

```bash
openclaw models set ollama/qwen3:8b
```

Verify the current default,

```bash
openclaw models
```

### START THE GATEWAY

Start the gateway service and verify it's running,

```bash
openclaw gateway start
openclaw health
systemctl --user status openclaw-gateway.service
```

First interaction with your AI assistant,

```bash
openclaw agent --agent main --message "say hello" --local
```

## ACCESS OPENCLAW

Two ways to access OpenClaw — terminal (TUI) or browser via SSH tunnel.

### TERMINAL (TUI)

```bash
openclaw tui
```

### BROWSER VIA SSH TUNNEL

OpenClaw's gateway runs over plain HTTP. To access it from
your workstation's browser safely, tunnel through SSH.

From the machine you'll use the browser on,

```bash
ssh -N -L 18789:127.0.0.1:18789 jeff@192.168.20.143
```

Then open in your browser (replace `<token>` with your gateway token),

```text
http://localhost:18789/#token=<token>
```

The gateway token was auto-generated during `openclaw onboard`.
Find it in your OpenClaw config,

```bash
grep -i token ~/.openclaw/openclaw.json
```
