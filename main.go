package main

import (
	"fmt"
	"time"

	"github.com/cdated/cycles/events"
)

func main() {
	fmt.Println("Create event")
	e1 := events.NewEventRecurs("Use device a", time.Second)
	fmt.Println(e1)

	fmt.Println("Complete event")
	e1.Complete()
	fmt.Println(e1)

	fmt.Println("Attempt reopen event")
	e1.Reopen()
	fmt.Println(e1)

	fmt.Println("Wait...attempt reopen event")
	time.Sleep(time.Second * 2)
	e1.Reopen()
	fmt.Println(e1)

	e2 := events.NewEvent("Use device b")
	fmt.Println(e2)
}
