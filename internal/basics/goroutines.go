package basics

import (
	"fmt"
	"strconv"
)

func printMessages(n int, message string) {
	for i := 0; i < n; i++ {
		fmt.Printf(message+"(%s)\n", strconv.Itoa(n))
	}
}

func TestPackage() {
	go printMessages(50, "Ping...")
	go printMessages(60, "Pong...")

	var input_temp string
	fmt.Scanln(&input_temp)
}
