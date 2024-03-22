package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func generateRandomNumbers(numChan chan<- int, step int) {
	defer wg.Done()
	for {
		numChan <- rand.Intn(50)
		time.Sleep(time.Second * time.Duration(step))
	}
}

func showNumbers(numChan <-chan int, step int) {
	defer wg.Done()
	for {
		randomNum := <-numChan
		for i := 0; i < step; i++ {
			fmt.Printf("%d <---\n", randomNum)
			time.Sleep(time.Second * 1)
		}
	}
}

func ExecuteRandomNumbersExample() {
	wg.Add(2)
	randomNumChan := make(chan int)
	step := 3
	go showNumbers(randomNumChan, step)
	go generateRandomNumbers(randomNumChan, step)

	wg.Wait()
}
