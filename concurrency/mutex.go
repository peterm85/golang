package main

import (
	"fmt"
	"sync"
)

var x = 0

func runWithRaceCondition() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(wg *sync.WaitGroup) {
			x = x + 1
			wg.Done()
		}(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

func runWithMutex() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func(wg *sync.WaitGroup, m *sync.Mutex) {
			m.Lock()
			x = x + 1
			m.Unlock()
			wg.Done()
		}(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

/*func main() {
	arg := os.Args[1]

	switch arg {
	case "1":
		runWithRaceCondition()
	case "2":
		runWithMutex()
	}
}*/
