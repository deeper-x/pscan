package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args[:]) != 2 {
		log.Println("missing destination")
		return
	}

	inaddr := os.Args[1]
	for i := 1; i <= 1024; i++ {
		res := scanport(inaddr, i)
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
