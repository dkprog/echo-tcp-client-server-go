package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:7000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected into", conn.RemoteAddr().String())
	defer conn.Close()
	codes := []string{"alpha", "bravo", "charlie", "delta", "echo"}
	for _, code := range codes {
		conn.SetDeadline(time.Now().Add(3 * time.Second))

		n, err := conn.Write([]byte(code))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(">>", code)

		buf := make([]byte, n)
		n, err = conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("<<", string(buf))
	}
}
