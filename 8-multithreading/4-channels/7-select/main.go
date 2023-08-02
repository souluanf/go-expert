package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{Id: i, Msg: "Hello from RabbitMQ"}
			c1 <- msg
		}

	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{Id: i, Msg: "Hello from Kafka"}
			c2 <- msg
		}
	}()

	//for i := 0; i < 3; i++ {
	for {
		select {
		case msg := <-c1: // RabbitMQ
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", msg.Id, msg.Msg)
		case msg := <-c2:
			fmt.Printf("Received from Kafka: ID: %d - %s\n", msg.Id, msg.Msg)
		case <-time.After(time.Second * 3):
			println("Timeout")
			//default:
			//	println("Default")
		}
	}
	//}
}
