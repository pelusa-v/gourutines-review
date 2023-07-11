package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func printMessages2(step int) {
	n := 10
	for i := 0; i < n; i++ {
		fmt.Printf("%d : %d\n", i, step)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Microsecond * amt)
	}
}

func TestDelayMessage() {
	for i := 0; i < 10; i++ {
		go printMessages2(i)
	}

	var input string
	fmt.Scanln(&input)
}
