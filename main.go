package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("socket and bind address required, for example:")
		fmt.Println(os.Args[0] + " /tmp/socket.sock 0.0.0.0:8080")
		return
	}
	socket := os.Args[1]
	listenAddr := os.Args[2]

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		panic(err)
	}

	log.Println("Waiting for connections on " + listenAddr + "...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept failed: %s", err.Error())
		}
		go handleConnection(conn, socket)
	}
}

func handleConnection(connTcp net.Conn, socket string) {
	connUnix, err := net.Dial("unix", socket)
	if err != nil {
		log.Printf("Failed to connect to socket \"%s\": %s", socket, err.Error())
	}

	log.Println("Accepted client")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer connUnix.Close()
		defer connTcp.Close()
		pipe(connUnix, connTcp)
	}()

	go func() {
		defer wg.Done()
		defer connUnix.Close()
		defer connTcp.Close()
		pipe(connTcp, connUnix)
	}()

	wg.Wait()

	log.Println("Finished with client")
}

func pipe(c1 net.Conn, c2 net.Conn) error {
	b := make([]byte, 8192)
	for {
		size, err := c1.Read(b)
		if err != nil {
			return err
		}
		_, err = c2.Write(b[:size])
		if err != nil {
			return err
		}
	}
}
