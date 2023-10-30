package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate random integers indefinitely
func generateRandInt(min, max int) <-chan int {
	// Create the output channel with a buffer 3, so there will only be 3 values at a time
	out := make(chan int, 3)

	// This go routine places data into the output channel to be read out of later
	go func() {
		for {
			out <- rand.Intn(max-min) + min
		}
	}()
	return out
}

// Generate random integers a set number of times
func generateRandIntn(count, min, max int) <-chan int {
	// Create the output channel with a buffer 3, so there will only be 3 values at a time
	out := make(chan int, 1)

	// This go routine places data into the output channel to be read out of later
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min) + min
		}
		close(out)
	}()
	return out
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// ***Use infinite generator***
	randInt := generateRandInt(1, 100)

	fmt.Println("generateRandInt infinite")

	// Read from the channel several times
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	// ***Use "n number of times" generator***
	randIntnRange := generateRandIntn(2, 1, 10)

	// when using a range loop on a channel, when the channel is closed, the loop will terminate
	// ...so this loop will end when randIntnRange is closed
	for i := range randIntnRange {
		fmt.Println("randIntnRange:", i)
	}

	// ***Use "n number of times" generator while listening for close***
	// (Essentially the same as above, but with more control)
	randIntnRange2 := generateRandIntn(4, 1, 10)

	for {
		n, open := <-randIntnRange2
		if !open {
			break
		}
		fmt.Println(n)
	}
}
