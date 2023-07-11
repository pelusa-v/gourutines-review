package channels

import (
	"fmt"
	"time"
)

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func TestChannel() {
	var c chan string = make(chan string)

	go printer(c)
	go pinger(c)
	go ponger(c)

	var input string
	fmt.Scanln(&input)
}
