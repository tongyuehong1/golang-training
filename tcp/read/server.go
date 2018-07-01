package main

import (
	"log"
	"net"
	"fmt"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		// 将read的数据放入到buf
		var buf = make([]byte, 10)
		log.Println("start to read from conn")


		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		fmt.Println("ddddd", buf)
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		log.Println("accept a new connection")
		go handleConn(c)
	}
}
