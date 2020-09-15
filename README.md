# TCP echo client/server in Go

## Purpose

As I'm learning Go, I found an [Echo client/server](https://en.wikipedia.org/wiki/Echo_Protocol) is the next logical step after messing around with basic TCP/IP [client](https://github.com/dkprog/simple-go-tcp-client) and [server](https://github.com/dkprog/simple-go-tcp-server) programming.

## Usage

### Running client

```
go run client/client.go
```

#### More options

```
go run client/client.go --help
```

### Running server

```
go run server/server.go
```

#### More options

```
go run server/server.go --help
```

## Troubleshooting

A typical TCP Echo server listens for connections on [TCP port 7](https://tools.ietf.org/html/rfc862). In this very example, I'm using `7000` so you don't have to run it as root. However, if you want to troubleshoot it against a real server, remember to change the bind/connect port using `--port` flag.

### Enable TCP Echo Protocol on inetd

Add the following lines to `/etc/inetd.conf` file:

```
echo   stream  tcp     nowait  root    internal
```

### Run client against inetd's Echo

```
go run client/client.go --port 7
```

### Test server using netcat

```
nc -v localhost 7000
```
