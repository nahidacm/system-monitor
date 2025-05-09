package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <PID | command name | username>")
		return
	}

	query := os.Args[1]
	allProcs, err := process.Processes()
	if err != nil {
		log.Fatalf("Failed to list processes: %v", err)
	}

	matchingProcs := []*process.Process{}

	// Try parsing as PID first
	if pid, err := strconv.Atoi(query); err == nil {
		for _, p := range allProcs {
			if int(p.Pid) == pid {
				matchingProcs = append(matchingProcs, p)
				break
			}
		}
	} else {
		// Otherwise check by name or username
		for _, p := range allProcs {
			name, _ := p.Name()
			username, _ := p.Username()

			if strings.EqualFold(name, query) || strings.EqualFold(username, query) {
				matchingProcs = append(matchingProcs, p)
			}
		}
	}

	if len(matchingProcs) == 0 {
		fmt.Println("No matching processes found.")
		return
	}

	for _, proc := range matchingProcs {
		name, _ := proc.Name()
		username, _ := proc.Username()
		cpuPercent, _ := proc.CPUPercent()
		memInfo, _ := proc.MemoryInfo()
		children, _ := proc.Children()

		fmt.Printf("----\nPID: %d\nName: %s\nUser: %s\n", proc.Pid, name, username)
		fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent)
		if memInfo != nil {
			fmt.Printf("Memory Usage: %.2f MB\n", float64(memInfo.RSS)/1024/1024)
		}
		fmt.Printf("Number of Subprocesses: %d\n", len(children))
	}
}
