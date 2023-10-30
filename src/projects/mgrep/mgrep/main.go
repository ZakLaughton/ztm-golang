package main

import (
	"fmt"
	"mgrep/worker"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexflint/go-arg"
)

// When it runs into  directory, recurse into it and pull out all the file paths
// All file paths will be added to the worklist
func discoverDirs(wl *worklist.Worklist, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Readdir error:", err)
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			nextPath := filepath.Join(path, entry.Name())
			discoverDirs(wl, nextPath)
		} else {
			wl.Add(worklist.NewJob(filepath.Join(path, entry.Name())))
		}
	}
}

// Commands from github.com/alexflint/go-arg
var args struct {
	SearchTerm string `arg:"positional, required"`
	SearchDir  string `arg:"positional"`
}

func main() {
	arg.MustParse(&args)

	// Create workgroup for workers
	var workersWg sync.WaitGroup

	// Create a buffer size of 100; 100 jobs are created before things start blocking
	wl := worklist.New(100)

	// Make a channel for results
	results := make(chan worker.Result, 100)

	numWorkers := 10

	workersWg.Add(1)

	go func() {
		defer workersWg.Done()
		discoverDirs(&wl, args.SearchDir)
		wl.Finalize(numWorkers)
	}()

	// spawn workers
	for i := 0; i < numWorkers; i++ {
		// Add 1 to the workersWg because we are starting a go routine
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			for {
				workEntry := wl.Next()
				if workEntry.Path != "" {
					workerResult := worker.FindInFile(workEntry.Path, args.SearchTerm)
					if workerResult != nil {
						for _, r := range workerResult.Inner {
							results <- r
						}
					}
				} else {
					return
				}
			}
		}()
	}

	// Create a channel and wrap up workersWg in it
	// Wait on workers to finish, but also display results while they are working
	blockWorkersWg := make(chan struct{})
	go func() {
		// sit and block, waiting for workers to finish
		// allows to select on other things while it's blocked
		workersWg.Wait()
		close(blockWorkersWg)
	}()

	// New waitgroup, specifically for displaying results
	var displayWg sync.WaitGroup

	displayWg.Add(1)
	go func() {
		for {
			select {
			case r := <-results:
				fmt.Printf("%v[%v]:%v\n", r.Path, r.LineNum, r.Line)
			case <-blockWorkersWg:
				// results may still be printing out
				if len(results) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()
	displayWg.Wait()
}
