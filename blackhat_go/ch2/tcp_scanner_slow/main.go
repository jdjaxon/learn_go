package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Port is closed or filtered.
			continue
		}
		conn.Close()
		fmt.Println(address, "open")
	}
}
