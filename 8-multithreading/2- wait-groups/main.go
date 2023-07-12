package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(100 * time.Millisecond)
		wg.Done() // Thread 2 and Thread 3 notify Thread 1 that they are done
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	// Thread 2
	go task("A", &waitGroup)
	// Thread 3
	go task("B", &waitGroup)
	// Thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(100 * time.Millisecond)
			waitGroup.Done() // Thread 4 notifies Thread 1 that it is done
		}
	}()
	waitGroup.Wait() // Thread 1 waits for all other threads to finish
}
