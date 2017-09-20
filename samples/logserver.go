package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8189")
	if err != nil {
		// handle error
	}

	fmt.Println("fmt:启动日志监听器, 8189端口")
	log.Println("log:启动日志监听器, 8189端口")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			fmt.Print("错误日志")
			// handle error
			continue
		}

		log.Print(conn)
		fmt.Print("普通日志")
		//go handleConnection(conn)
	}
}
