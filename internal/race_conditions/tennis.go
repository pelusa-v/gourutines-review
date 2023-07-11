package race_conditions

import (
	"fmt"
	"math/rand"
	"time"
)

func ExecuteGame() {
	game := make(chan int)

	wg.Add(2)

	go play("Del Potro", game)
	go play("Federer", game)

	game <- 1

	wg.Wait()
}

func play(name string, step chan int) {
	defer wg.Done()

	for {
		time.Sleep(time.Millisecond * 500)
		ball, ok := <-step

		if !ok {
			fmt.Printf("%s won (channel closed)\n", name)
			return
		}

		if rand.Intn(100)%13 == 0 {
			fmt.Printf("%s missed!\n", name)
			close(step)
			return
		}

		ball++
		fmt.Printf("%s hit the ball %d\n", name, ball)
		step <- ball
	}
}
