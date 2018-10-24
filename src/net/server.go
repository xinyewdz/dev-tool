// server.go
package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()
			var (
				n   int
				b   = make([]byte, 1024)
				err error
			)
			n, err = conn.Read(b)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(string(b[:n]))
			n, err = conn.Write([]byte("pong"))
			if err != nil {
				log.Println(err)
				return
			}
		}(conn)
	}
}
