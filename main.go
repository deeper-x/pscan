package main

import (
	"log"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "192.168.1.1:80")
	if err != nil {
		log.Println(err)
		log.Println("closed")
		return
	}

	log.Println("open")
}
