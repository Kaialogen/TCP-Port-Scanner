//--------------------------------
// GO TCP SCANNER v0.0.5-alpha
//
// Made by Kaialogen
//
//--------------------------------

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
	"time"
)

func worker(ports, results chan int, ip string, portType string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", ip, p)
		conn, err := net.Dial(portType, address)
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
	portType := flag.String("type", "tcp", "Type of port to scan (tcp)")
	flag.Parse()
	// check for correct usage
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: ", os.Args[0], "[ip] -type [tcp]")
		os.Exit(1)
	}

	ip := flag.Args()[0]

	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	start := time.Now()

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, ip, *portType)
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
