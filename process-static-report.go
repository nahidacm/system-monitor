package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <PID>")
		return
	}

	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid PID: %v", err)
	}

	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		log.Fatalf("Failed to find process with PID %d: %v", pid, err)
	}

	// CPU Percent
	cpuPercent, err := proc.CPUPercent()
	if err != nil {
		log.Printf("Failed to get CPU percent: %v", err)
	} else {
		fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent)
	}

	// Memory Info
	memInfo, err := proc.MemoryInfo()
	if err != nil {
		log.Printf("Failed to get memory info: %v", err)
	} else {
		fmt.Printf("Memory Usage: %d bytes\n", memInfo.RSS)
	}

	// Number of subprocesses
	children, err := proc.Children()
	if err != nil {
		log.Printf("Failed to get child processes: %v", err)
	} else {
		fmt.Printf("Number of Subprocesses: %d\n", len(children))
	}
}
