package logical_processors

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func printNumbers(id string, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		fmt.Printf("%s : %d\n", id, i)
	}
}

func ExecuteConcurrentPrinter() {
	// runtime.GOMAXPROCS(1)
	// runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(3)

	go printNumbers("A", 5000)
	go printNumbers("B", 5000)
	go printNumbers("C", 5000)

	fmt.Printf("%s\n", "Waiting for finish...")
	wg.Wait()
	fmt.Printf("Logical Processors: %d\n", runtime.NumCPU())
}
