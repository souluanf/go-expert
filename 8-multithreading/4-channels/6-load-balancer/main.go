package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data <-chan int) {
	for d := range data {
		fmt.Printf("Worker %d received %d\n", workerID, d)
		time.Sleep(time.Second)
	}

}

func main() {
	data := make(chan int)
	qtdWorkers := 1000000

	// Create workers
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}
	for i := 0; i <= 10000000; i++ {
		data <- i
	}
}
