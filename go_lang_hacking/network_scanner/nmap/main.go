package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap"
)

func main() {
	targetIP := "scanme.nmap.org"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	scanner, err := nmap.NewScanner(
		nmap.WithTargets(targetIP),
		nmap.WithPorts("80", "443"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("error creating scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if warnings != nil {
		log.Printf("warnings: \n%v", warnings)
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])
		if len(host.Addresses) > 1 {
			fmt.Printf("  - MAC addresses: %v\n", host.Addresses[1:])
		}

		for _, port := range host.Ports {
			fmt.Printf("  - Port %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}
}