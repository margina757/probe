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

	checkOS()
	fmt.Println("\nLonlife Network Delay Probe Program\n")
	db.OpenDB()
	probe.Start()
}

func checkOS() {
	if runtime.GOOS != "linux" {
		fmt.Println("Linux Support Only")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}
}
