package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func JobWithCtx(ctx context.Context, jobID int) error {
	select {
	case <-ctx.Done(): //cancel when Ctrl + C
		fmt.Printf("context cancelled job %v terminating\n", jobID)
		return nil
	case <-time.After(time.Second * time.Duration(rand.Intn(3))): //sleep
	}
	if rand.Intn(12) == jobID {
		fmt.Printf("Job %v failed.\n", jobID)
		return fmt.Errorf("job %v failed", jobID)
	}

	fmt.Printf("Job %v done.\n", jobID)
	return nil
}

func cancellableCtx() (*errgroup.Group, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	go func() {
		sCh := make(chan os.Signal, 1)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM) // Ctrl + C
		<-sCh
		cancel()
	}()
	return eg, ctx
}

func runJobs(eg *errgroup.Group, ctx context.Context) {
	for i := 0; i < 10; i++ {
		jobID := i
		eg.Go(func() error {
			return JobWithCtx(ctx, jobID)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("Encountered error:", err)
	}
	fmt.Println("Successfully finished.")
}

/*func main() {
	arg := os.Args[1]

	var eg *errgroup.Group
	var ctx context.Context

	switch arg {
	case "1":
		eg, ctx = errgroup.WithContext(context.Background())
	case "2":
		eg, ctx = cancellableCtx()
	}

	runJobs(eg, ctx)
}*/
