package main

import (
	"fmt"
	"time"
)

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)

    b := <-table
	fmt.Println("total hits", b.Hits)
}

type Ball struct {
	Hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.Hits++
		fmt.Println(name, ball.Hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}