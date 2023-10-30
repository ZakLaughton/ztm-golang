package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOk
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		// Select will continuously check for messages on the channels control and jobs
		select {
		case msg := <-control:
			switch msg {
			// Exit goroutine when DoExit is received
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOk
				return
			// Catch unhandled exceptions
			default:
				panic("unhandled control message")
			}
		case job := <-jobs:
			// Create a new result and send it to the results channel
			results <- Result{result: job.data * 2, job: job}
		}
		// Once done, go back to the top of the loop and check for messages in both channels again
	}
}

func main() {
	jobs := make(chan Job, 50)
	results := make(chan Result, 50)
	control := make(chan ControlMsg)

	go doubler(jobs, results, control)

	// Send 30 jobs to the jobs channel for the go routine above to read
	// doubler will send results to the result channel
	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			// Send DoExit to control channel
			control <- DoExit
			// Wait for control channel to respond
			<-control
			fmt.Println("program exit")
			return
		}
	}
}
