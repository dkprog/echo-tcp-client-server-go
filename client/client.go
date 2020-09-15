package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	host := flag.String("host", "localhost", "The Echo server's hostname")
	port := flag.Int("port", 7000, "The Echo server's TCP port")

	flag.Parse()

	address := fmt.Sprintf("%s:%d", *host, *port)

	conn, err := net.Dial("tcp", address)
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
