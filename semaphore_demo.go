package main

import (
	j "play-with-concurrency/java_util_concurrent"
	"fmt"
	"sync"
)

var sharedCount int32

func main() {
	sharedCount = 0

	s := j.NewSemaphore()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go intrementer(s, wg)
	go decrementer(s, wg)

    wg.Wait()
	fmt.Println(sharedCount)
}

func intrementer(s *j.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	s.AcquireN(10)
	for i := 0; i < 1000; i++ {
		sharedCount++
	}
	s.ReleaseN(2)
	s.ReleaseN(8)
	fmt.Println("exit incr")
}

func decrementer(s *j.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	s.AcquireN(5)
	for i := 0; i < 1000; i++ {
		sharedCount--
	}
	s.ReleaseN(5)

	fmt.Println("exit decr")
}