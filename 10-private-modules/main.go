package main

import (
	"fmt"
	"github.com/souluanf/lfutils-secret-go/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
