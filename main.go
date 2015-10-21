package main

import (
	"fmt"
	"os"
	"probe/db"
	"probe/probe"
	"runtime"
	"time"
)

func main() {
	db.OpenDB()
	checkOS()
	fmt.Println("\nLonlife Network Delay Probe Program\n")
	probe.Start()
}

func checkOS() {
	if runtime.GOOS != "linux" {
		fmt.Println("Linux Support Only")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}
}
