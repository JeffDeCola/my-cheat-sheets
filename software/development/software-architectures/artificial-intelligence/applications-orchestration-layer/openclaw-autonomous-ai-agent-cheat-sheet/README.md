# OPENCLAW AUTONOMOUS AI-AGENT CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_OpenClaw is a personal AI assistant that runs on your server,
connects to your apps, and does things for you automatically._

tl;dr

```text
openclaw tui                  # The main text interface
```

Table of contents

* tbd

## OVERVIEW

The AI Software/Hardware Stack

```text
      OpenClaw                  The orchestrator - listens, acts, manages (not AI)
      Ollama                    The inference server - hosts and serves the AI model
      llama, qwen, deepseek     The AI model — The frozen brain
      P40                       The hardware accelerator - for the frozen brain
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

WHAT OPENCLAW ACTUALLY DOES

```bash text
Feeds knowledge    — pulls real-time context (weather, files, web)
Takes actions      — sends emails, creates events, runs code
Remembers          — persistent memory across sessions
Plans              — breaks tasks into steps and executes them
Runs autonomously  — wakes up on a schedule without you asking
```

## INSTALL

This is how I installed OpenClaw on my dell rack server that has
a nvidia Tesla P40 in it.

OpenClaw installs via Node.js and npm on Ubuntu.

```bash
sudo apt update && sudo apt upgrade -y
```

Install Node.js v22 via NVM

```bash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
source ~/.bashrc
nvm install 22
nvm use 22
node --version
```

Install openclaw via node package manager

```bash
npm install -g openclaw
```

Verify

```bash
openclaw --version
```

Fix Ollama to listen on all interfaces (run on VM 102)

```bash
sudo systemctl edit ollama

```

Add these lines:

```text
[Service]
Environment="OLLAMA_HOST=0.0.0.0:11434"
```

```bash
sudo systemctl daemon-reload
sudo systemctl restart ollama
ss -tlnp | grep 11434    # should show *:11434
```

Verify Ollama is reachabl

```bash
curl http://192.168.20.141:11434
# Should return: Ollama is running
```

Start openclaw onboard, the interactive setup wizard that configures OpenClaw

```bash
openclaw onboard
```

Onboarding selections:

```text
Security warning       → Yes, continue
Gateway                → Local gateway (this machine)
Workspace              → default (~/.openclaw/workspace)
Model provider         → Ollama
Ollama base URL        → http://192.168.20.141:11434
Ollama mode            → Local
Default model          → ollama/qwen2.5-coder:32b   # you can change this later
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

SEt the correct model

```bash
openclaw models set ollama/qwen2.5-coder:32b
```

Start the gateway

```bash
openclaw gateway start
openclaw health
systemctl --user status openclaw-gateway.service
```

Your first interaction with you AI assistant.

```bash
openclaw agent --agent main --message "say hello" --local
```

## ACCESS OPENCLAW

You can access openclaw using a terminal or a browser.

### TERMINAL

For the terminal its just

```bash
openclaw tui
```

### BROWSER

VIA BROWSER
For the Browser, since openclaw is http, we need to do one of two things

* A SSH Tunnel
* A reverse proxy with https

#### VIA SSH TUNNEL

Setup an ssh tunnel on the machine you're using the browser from

```bash
ssh -N -L 18789:127.0.0.1:18789 jeff@192.168.20.143
```

[http://localhost:18789/#token=xxxx]()

#### VIA REVERSE PROXY (COULD NOT GET TO WORK)

The flow

```text
Browser (Windows)
   ↓
NGINX (192.168.20.161, LXC #201)
   ↓
OpenClaw Agent (192.168.20.143, VM #104)
```

Create a lxc that has nginx

On the lxc nginx server create a certificate

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout /etc/ssl/private/nginx-selfsigned.key \
  -out /etc/ssl/certs/nginx-selfsigned.crt \
  -subj "/CN=homelab/O=homelab/C=US"
```

Now create a new site cert for your openclaw server

```bash
nano /etc/nginx/sites-available/openclaw-jeff
```

```text
server {
    listen 443 ssl;
    server_name _;

    # SSL certificate and key
    ssl_certificate /etc/ssl/certs/nginx-selfsigned.crt;
    ssl_certificate_key /etc/ssl/private/nginx-selfsigned.key;

    # Optional SSL settings for better security
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers on;

    # Main location block for OpenClaw
    location / {
        proxy_pass http://192.168.20.143:18789;
        proxy_http_version 1.1;

        # WebSocket headers
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Origin $scheme://$host;

        # Standard proxy headers
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # Optional: WebSocket timeout settings
        proxy_read_timeout 3600s;
        proxy_send_timeout 3600s;
    }
}
```

Currently in /etc/nginx/sites-available

```bash
lrwxrwxrwx 1 root root   34 Mar 22 16:54 default -> /etc/nginx/sites-available/default
```

Lets change that

```bash
ln -s /etc/nginx/sites-available/openclaw-jeff /etc/nginx/sites-enabled/
rm /etc/nginx/sites-enabled/default
```

We now have

```bash
lrwxrwxrwx 1 root root   40 Mar 22 17:01 openclaw-jeff -> /etc/nginx/sites-available/openclaw-jeff
```

test its working

```bash
nginx -t
```

And reload it

```bash
systemctl reload nginx
```

Now add nginx to openclaw config file on your openclaw server

```bash
nano ~/.openclaw/openclaw.json
```

```json
  "allowedOrigins": [
  "http://localhost:18789",
  "http://127.0.0.1:18789",
  "http://192.168.20.143:18789",
  "https://192.168.20.161"
]
```

And add to gateway section

```json
"trustedProxies": ["192.168.20.161/32"],
```

```bash
systemctl --user restart openclaw-gateway.service && systemctl --user restart openclaw-node.service
```

```text
Your browser → https://192.168.20.161 → Nginx → OpenClaw
```
