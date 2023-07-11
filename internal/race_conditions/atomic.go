package race_conditions

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	SHUTDOWN int64
)

func SendDataToServer(data string) {
	defer wg.Done()

	for {
		fmt.Printf("%s : %s\n", "Sending data to server...", data)
		time.Sleep(300 * time.Millisecond)

		if atomic.LoadInt64(&SHUTDOWN) == 1 {
			fmt.Printf("The server is shutdown, data have not sent")
			break
		}
	}
}

func ConfigureServer() {
	wg.Add(2)

	go SendDataToServer("IP LINUX")
	go SendDataToServer("IP WINDOWS")

	time.Sleep(10 * time.Second)

	atomic.StoreInt64(&SHUTDOWN, 1)
	fmt.Printf("%s\n", "Server SHUTDOWN")

	wg.Wait()
}
