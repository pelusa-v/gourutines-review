package logical_processors

import (
	"fmt"
)

func printUppercaseAlphabet(n int) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		for letter := 'A'; letter < 'A'+26; letter++ {
			fmt.Printf("%c ", letter)
		}
	}
}

func printLowercaseAlphabet(n int) {
	defer wg.Done()

	for i := 0; i < n; i++ {
		for letter := 'a'; letter < 'a'+26; letter++ {
			fmt.Printf("%c ", letter)
		}
	}
}

func ExecuteAlphabetPrinter() {
	// runtime.GOMAXPROCS(1)
	// runtime.GOMAXPROCS(runtime.NumCPU()) // default
	wg.Add(2)

	go printLowercaseAlphabet(3)
	go printUppercaseAlphabet(3)

	fmt.Println("Waiting for finish...")
	wg.Wait()
	fmt.Println()
}
