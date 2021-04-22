package main

import (
	"fmt"
	"time"
)

func runGoroutines() {
	for i := 0; i < 10; i++ {
		go fmt.Printf("This is job: %v\n", i)
	}
	time.Sleep(2 * time.Second)
}

/*func main() {
	runGoroutines()
}*/
