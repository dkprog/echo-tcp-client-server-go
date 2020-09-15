package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.Int("port", 7000, "This Echo server's TCP port to listen to")
	flag.Parse()
	address := fmt.Sprintf(":%d", *port)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening for new connections on %s", address)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var remote string = conn.RemoteAddr().String()
	log.Printf("%s connected.", remote)

	defer conn.Close()

	for {
		conn.SetDeadline(time.Now().Add(3 * time.Second))

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Print(err)
			break
		} else {
			log.Printf("Received %d bytes from %s.\n", n, remote)
		}
		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Print(err)
			break
		} else {
			log.Printf("Sent %d bytes from %s.\n", n, remote)
		}
	}
}
