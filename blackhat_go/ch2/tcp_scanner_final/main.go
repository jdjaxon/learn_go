package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan string) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- ""
			continue
		}

		conn.Close()
		results <- address
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan string)
	var openports []string

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		address := <-results
		if address != "" {
			openports = append(openports, address)
		}
	}

	close(ports)
	close(results)
	sort.Strings(openports)
	for _, address := range openports {
		fmt.Println(address, "open")
	}
}
