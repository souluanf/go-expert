package main

import (
	"fmt"
	"time"
)

// Thread 1
func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
		forever <- true
	}()
	<-forever
}
