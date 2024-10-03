package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	IP := "scanme.nmap.org"
	address := fmt.Sprintf(IP+":%d", port)
	
	// Adding a timeout for connection attempts to avoid hanging
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	
	if err == nil {
		fmt.Printf("[+] Connection Established... PORT %v\n", port)
		// Close the connection immediately after establishing it
		conn.Close()
	} else {
		fmt.Printf("[-] Connection Failed on PORT %v\n", port)
	}
}

func main() {

	var wg sync.WaitGroup
	// Scan the first 1024 ports
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go scanPort(i, &wg)
	}
	wg.Wait()
}
