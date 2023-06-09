//--------------------------------
// GO TCP/UDP SCANNER v0.0.2-alpha
//
// Made by Kaialogen
//
//--------------------------------

package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"time"
)

func worker(ports, results chan int, ip_input, port_type string) {
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

	// check for correct usage
	if len(os.Args) != 3 {
		fmt.Println("Usage: ", os.Args[0], "[ip] [tcp/udp]")
		os.Exit(1)
	}

	ip_input := os.Args[1]
	port_type := os.Args[2]

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	start := time.Now()

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, ip_input, port_type)
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
