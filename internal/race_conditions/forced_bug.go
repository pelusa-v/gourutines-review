package race_conditions

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	COUNTER int
	wg      sync.WaitGroup
)

func TestForcedRaceCondition() {
	wg.Add(2)

	go increaseCounter("A", 2)
	go increaseCounter("B", 2)

	fmt.Printf("%s\n", "Wait for finish...")
	wg.Wait()
	fmt.Printf("Final value of COUNTER: %d\n", COUNTER)
}

func increaseCounter(id string, n int) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		localCounter := COUNTER
		runtime.Gosched() // To force race condition
		localCounter++
		COUNTER = localCounter
	}
}
