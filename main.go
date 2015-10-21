package main

import (
	"fmt"
	"os"
	"probe/probe"
	"runtime"
)

func main() {
	checkOS()
	fmt.Println("\nLonlife Network Delay Probe Program\n")
	probe.Start()
}

func checkOS() {
	if runtime.GOOS != "linux" {
		fmt.Println("Linux Support Only")
		os.Exit(0)
	}
}
