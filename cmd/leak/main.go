package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	for i := range 1000 {
		isq, err := square(ctx, i)
		if err == nil {
			fmt.Printf("%d ^ 2 = %d\n", i, isq)
		} else {
			fmt.Printf("%d ^ 2 failed: %v\n", i, err)
		}
	}

	time.Sleep(1 * time.Second)

	runtime.GC()
	<-exit
	runtime.GC()
	cancel()
}

func square(ctx context.Context, input int) (int, error) {
	ctxT, _ := context.WithCancel(ctx)

	result := make(chan int)
	var err error

	go func() {
		for {
			select {
			case <-time.After(1 * time.Millisecond):
				result <- input * input
			case <-ctxT.Done():
				err = ctxT.Err()
				result <- -1
			}
		}
	}()

	return <-result, err
}
