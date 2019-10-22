package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	// memStats := new(runtime.MemStats)
	// runtime.ReadMemStats(memStats)
	fmt.Println("arc", runtime.GOARCH)
	fmt.Println("os", runtime.GOOS)
	fmt.Println("cpus", strconv.Itoa(runtime.NumCPU()))
}
