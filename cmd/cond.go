package main

import (
	"fmt"
	"sync"
	"time"
)

var stop int

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	go worker("One", cond)
	go worker("Two", cond)
	go worker("Three", cond)
	go worker("Four", cond)

	cond.Broadcast()
	cond.Broadcast()
	cond.Broadcast()
	cond.Broadcast()
	cond.Broadcast()
	cond.Broadcast()

	cond.Broadcast()

	stop = 1

	cond.Broadcast()

	time.Sleep(5 * time.Second)

	fmt.Println("Done")

}

func worker(name string, cond *sync.Cond) {
	for {
		fmt.Println(name, "waits")
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()
		fmt.Println(name, "does work")

		if stop == 1 {
			fmt.Println(name, "stop")
			return
		}
	}
}
