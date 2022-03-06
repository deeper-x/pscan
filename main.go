package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for i := 1; i <= 100; i++ {
		res := scanport("192.168.1.1", i)
		if res {
			log.Printf("%d: open", i)
		}
	}
}

func scanport(addr string, p int) bool {
	addrport := fmt.Sprintf("%s:%d", addr, p)

	_, err := net.Dial("tcp", addrport)

	return err == nil
}
