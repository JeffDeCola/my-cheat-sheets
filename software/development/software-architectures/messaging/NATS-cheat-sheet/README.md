# NATS CHEAT SHEET

`NATS` _is a pub/sub messaging bus perfect for asynchronous microservice communication._

tl;dr,

```bash
nats-server -v
nats-server -DV -p 4222 -a 127.0.0.1
# INSTALL NATS GO CLIENT LIBRARIES
go get -v -u github.com/nats-io/nats.go/
```

Table of Contents,

*
*

Documentation and reference,

* [Official NATS site](https://nats.io/)
* [Official github page](https://github.com/nats-io/nats-server)
* For some examples with go, refer to my repo
  [my-go-examples](https://github.com/JeffDeCola/my-go-examples#messaging)

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## NATS

NATS is a Message BUS.
Fundamentally NATS is about publishing and listening for messages.
The official NATS server was `gnatsd`
but renamed to `nats-server` for the v2.0.0.release.

## INSTALL A NATS SERVER (nats-server)

There are many ways to install,

1. NATS Docker image
1. Get the binaries
1. Use the public NATS server nats://demo.nats.io:4222
1. Build from source
1. Package Manager

MacOS install,

```bash
brew install nats-server
```

Linux Install release build,

```bash
NATS_VER=2.0.4
curl -L https://github.com/nats-io/nats-server/releases/download/v$NATS_VER/nats-server-v$NATS_VER-linux-amd64.zip -o nats-server.zip
unzip nats-server.zip -d nats-server
sudo cp nats-server/nats-server-v$NATS_VER-linux-amd64/nats-server /usr/local/bin/.
rm nats-server.zip
rm -r nats-server
```

Raspberry pi, hummingboard, build from source,

```bash
GO111MODULE=on go get github.com/nats-io/nats-server/v2
```

check,

```bash
nats-server -v
```

## START NATS (nat-server)

For the default of `0.0.0.0:4222` (security is disabled),

```bash
nats-server
nats-server -DV -p 4222 -a 127.0.0.1
```

Where -DV is both debug and trace log.

The NATS server listens for client connections on TCP Port 4222.

## INSTALL NATS GO CLIENT LIBRARY

```go
go get -v -u github.com/nats-io/nats.go/
```

Refer to my repo
[my-go-examples](https://github.com/JeffDeCola/my-go-examples#messaging)
for how to connect to NATS and use via go.
