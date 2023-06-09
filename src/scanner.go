//--------------------------------
// GO TCP/UDP SCANNER v0.0.1
//
// Made by Kaialogen
//
//--------------------------------

package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

var ip_input = ""
var port_type = ""

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", ip_input, p)
		conn, err := net.Dial(port_type, address)
		if err != nil {
			//port is closed or filtered
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	fmt.Println("Enter IP to be scanned: ")
	fmt.Scan(&ip_input)

	fmt.Println("tcp or udp: ")
	fmt.Scan(&port_type)

	start := time.Now()

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

	elapsed := time.Since(start)
	fmt.Printf("Scan Complete. Time taken: %s\n", elapsed)
}
