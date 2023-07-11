package race_conditions

import (
	"fmt"
	"time"
)

var RUNNERS = [4]string{"Tom", "Paolo", "Joe", "Don"}

func ExecuteRace() {
	wg.Add(1)
	race_step := make(chan int)

	go ReceiveBaton(race_step)
	race_step <- 0

	wg.Wait()
}

func ReceiveBaton(race_step chan int) {

	baton_step, _ := <-race_step

	fmt.Println("--------------")
	if baton_step < 4 {
		fmt.Printf("%s está corriendo y lleva la posta %d\n", RUNNERS[baton_step], baton_step)
		time.Sleep(time.Millisecond * 1200)
		if baton_step != 3 {
			fmt.Printf("%s entrega la posta al siguiente corredor\n", RUNNERS[baton_step])
		}
		baton_step++
		go ReceiveBaton(race_step)
	}

	if baton_step == 4 {
		fmt.Println("El último corredor llegó a la meta")
		close(race_step)
		wg.Done()
		return
	}

	race_step <- baton_step
}
