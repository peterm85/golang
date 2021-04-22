package main

import (
	"fmt"
	"math/rand"
	"sync"

	"golang.org/x/sync/errgroup"
)

func runWaitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(jobID int) {
			defer wg.Done()
			fmt.Printf("This is job: %v\n", jobID)
		}(i)
	}
	wg.Wait()
}

func runErrGroup() {
	var eg errgroup.Group

	for i := 0; i < 10; i++ {
		jobID := i
		eg.Go(func() error {
			if rand.Intn(12) == jobID {
				return fmt.Errorf("job %v failed", jobID)
			} else {
				fmt.Printf("Job %v done.\n", jobID)
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("Encountered error:", err)
	}
	fmt.Println("Successfully finished.")
}

/*func main() {
	arg := os.Args[1]

	switch arg {
	case "1":
		runWaitGroup()
	case "2":
		runErrGroup()
	}
}*/
