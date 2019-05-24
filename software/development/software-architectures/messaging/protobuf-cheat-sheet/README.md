# PROTOBUF CHEAT SHEET

`protobuf` _short for protocol buffers is a method of
serializing structured data (useful for messaging)._

tl;dr,

```bash
protoc --version
protoc --go_out=. messages.proto
```

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

Define the message formats in a .proto file. The concept
is to have a human readable file representing your data. 

Then your client and your server can take that .proto
file and compile it (using protoc) to their language of
choice and use it in their microservice/app/service.

Here is an illustration,

![IMAGE - protobuf - IMAGE](../../../../../docs/pics/protobuf.jpg)

## INSTALL PROTOC FROM SOURCE

For macOS,

```bash
PROTOC_ZIP=protoc-3.7.1-osx-x86_64.zip
curl -OL https://github.com/google/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
rm -f $PROTOC_ZIP
```

For linux,

```bash
PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
curl -OL https://github.com/google/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
rm -f $PROTOC_ZIP
```

Check version,

```bash
protoc --version
```

## INSTALL LIBRARIES FOR GO

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

This will get the libraries and
place `protoc-gen-go` in `$GOPATH/bin/protoc-gen-go`.

## .PROTO FILE (VER3)

The .proto file defines the message format
in readable form.

As an example I use in my repo
[my-go-examples](https://github.com/JeffDeCola/my-go-examples/tree/master/messaging/protobuf)

```txt   
syntax = "proto3";

package main;

message Person {
	string name = 1;
    int32 age = 2;
	string email = 3;
    string phone = 4;
}

message AddressBook {
  repeated Person people = 1;
}
```

## RUN PROTOC

`protoc` is the protocol buffer compiler that will
compile the `messages.proto` file to a readable form
for your particular language.

```bash
protoc --go_out=. messages.proto
```

## USEFUL FOR RPC

Protobuf serves as a basis for a custom remote procedure
call (RPC) system that can be used for
inter-machine communication.

The whole purpose of using protocol buffers is to serialize
your data so that it can be parsed elsewhere.