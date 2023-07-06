package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(port int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println(address, "open")
		}(i)
	}
}
