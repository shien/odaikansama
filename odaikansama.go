package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid number of args.")
		os.Exit(1)
	}

	if os.Args[1] == "--apikey" {
		// start odaikan bot
		if len(os.Args[2]) > 0 {
			apikey := os.Args[2]
			Run(apikey)
		}

	} else {
		fmt.Println("Invalid option :" + os.Args[1])
	}
}
