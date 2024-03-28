package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // Get command-line arguments, excluding the path to the program

	if len(args) == 0 {
		fmt.Println("No command specified")
		return
	}

	switch args[0] {
	case "server1":
		server1()
	case "server2":
		server2()
	case "basicReserverProxy":
		basicReserverProxy()
	case "loadBalance":
		loadBalance()
	case "basicForwardProxy":
		basicForwardProxy()
	}
}
