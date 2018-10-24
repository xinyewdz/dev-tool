// client.go
package main

import (
	"log"
	"net"
)

func main() {
	var (
		conn net.Conn
		b    = make([]byte, 1024)
		n    int
		err  error
	)

	conn, err = net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("ping"))
	if err != nil {
		log.Fatalln(err)
	}
	n, err = conn.Read(b)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b[:n]))
}
