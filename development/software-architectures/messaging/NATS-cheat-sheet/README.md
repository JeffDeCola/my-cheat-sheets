# NATS CHEAT SHEET

`NATS` _is a messaging bus perfect for asynchronous microservice communication._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## NATS

NATS is a Message BUS.

## NATS SERVER

To use NATS you can do one of the following things:

1. Use the NATS Docker image
1. Get the binaries
1. Use the public NATS server nats://demo.nats.io:4222
1. Build from source

To install for mac,

```bash
brew install gnatsd
```

To start,

```bash
gnatsd
gnatsd --addr 127.0.0.1 --port 4222
```

## GO LIBRARY

```go
go get https://github.com/nats-io/nats
```
