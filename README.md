Glues a unix socket to a TCP socket.

Usage:

```
unix-to-tcp /tmp/unicorn.sock 0.0.0.0:8080
```

Building:

```
go get github.com/porty/unix-to-tcp
unix-to-tcp
```

Or if you want to cross compile and stuff:

```
GOOS=linux go build -o linux-unix-to-tcp github.com/porty/unix-to-tcp
```
