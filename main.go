package main

import (
	"fmt"
	"yoshi/animServer/deque"
)

players map [int64] *PlayerState


func main() {

	// For incoming commands
	DQin  := deque.DeqRing[Command]{}

	// For processed events
	DQout := deque.DeqRing[Event]  {}

	DQin.Initialize(1024 * 1024 * 8)
	DQout.Initialize(1024 * 1024 * 8)

	
	fmt.Println("Great job")
	return
}
