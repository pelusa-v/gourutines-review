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

	go increaseCounter(1)
	go increaseCounter(2)

	wg.Wait()
	// fmt.Printf("%s\n", "Wait for finish...")
	fmt.Printf("Final value of COUNTER: %d\n", COUNTER)
}

func increaseCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		localCounter := COUNTER
		runtime.Gosched() // To force race condition
		localCounter++
		COUNTER = localCounter
	}
}
