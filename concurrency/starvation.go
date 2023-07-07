package main

import (
	"fmt"
	"sync"
	"time"
)

// Define a WaitGroup to manage synchronization
var wg sync.WaitGroup

// Define a shared Mutex to manage concurrent access
var sharedLock sync.Mutex

// Define the total runtime for the workers
const runtime = 1 * time.Second

func main() {
	// Add 2 to the WaitGroup counter, since we are launching two goroutines
	wg.Add(2)
	// Launch a goroutine for the greedyWorker function
	go greedyWorker()
	// Launch a goroutine for the politeWorker function
	go politeWorker()

	// Wait for both goroutines to finish
	wg.Wait()
}

func politeWorker() {
	// Ensure the WaitGroup counter is decremented when the function finishes
	defer wg.Done()
	var count int
	// Loop for the duration of the runtime
	for begin := time.Now(); time.Since(begin) <= runtime; {
		// Lock the sharedLock Mutex to begin a unit of work
		sharedLock.Lock()
		// Simulate work by sleeping for a nanosecond
		time.Sleep(1 * time.Nanosecond)
		// Unlock the sharedLock Mutex to end the unit of work
		sharedLock.Unlock()

		// The following two blocks repeat the same pattern, simulating a polite worker
		// that only takes the lock for short periods of time      sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		// Increment the count of work units completed
		count++
	}
	// Print the count of work units completed by the polite worker
	fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
}

func greedyWorker() {
	// Ensure the WaitGroup counter is decremented when the function finishes
	defer wg.Done()
	var count int
	// Loop for the duration of the runtime
	for begin := time.Now(); time.Since(begin) <= runtime; {
		// Lock the sharedLock Mutex to begin a unit of work
		sharedLock.Lock()
		// Simulate work by sleeping for 3 nanoseconds, which is longer than the polite worker
		time.Sleep(3 * time.Nanosecond)
		// Unlock the sharedLock Mutex to end the unit of work
		sharedLock.Unlock()
		// Increment the count of work units completed
		count++
	}
	// Print the count of work units completed by the greedy worker
	fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
}
