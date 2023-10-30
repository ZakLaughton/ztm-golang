package main

import (
	"context"
	"fmt"
	"time"
)

// Context is useful for when you want to cancel a long running process without
// having to deal with channels and goroutines.

func sampleOperation(ctx context.Context, msg string, msDelay time.Duration) <-chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			// Emit message after delay
			case <-time.After(msDelay * time.Millisecond):
				out <- fmt.Sprintf("%v operation completed", msg)
				return
			// Abort if context is cancelled
			case <-ctx.Done():
				out <- fmt.Sprintf("%v aborted", msg)
			}
		}
	}()
	return out
}

func main() {
	// Create a new empty context
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)

	webServer := sampleOperation(ctx, "webserver", 200)
	microservice := sampleOperation(ctx, "microservice", 200)
	database := sampleOperation(ctx, "database", 200)

MainLoop:
	for {
		select {
		case val := <-webServer:
			fmt.Println(val)
		case val := <-microservice:
			fmt.Println(val)
			fmt.Println("cancel context")
			cancelCtx()
			break MainLoop
		case val := <-database:
			fmt.Println(val)
		}

		// Just to see what happens when we cancel the context to see if anything comes back from the database
		// We shouldn't, because the database takes the longest to run
		fmt.Println(<-database)
	}
}
