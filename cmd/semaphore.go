package main

import (
	"fmt"
	juc "github.com/ksopin/play-with-concurrency/pkg/java_util_concurrent"
	"sync"
)

var sharedCount int32

func main() {
	sharedCount = 0

	s := juc.NewSemaphore()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go incrementer(s, wg)
	go decrementer(s, wg)

	wg.Wait()
	fmt.Println(sharedCount)
}

func incrementer(s *juc.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	s.AcquireN(10)

	for i := 0; i < 10; i++ {
		for j := 0; j < 100; j++ {
			sharedCount++
		}
		s.ReleaseN(1)
	}
	fmt.Println("exit incr")
}

func decrementer(s *juc.Semaphore, wg *sync.WaitGroup) {
	defer wg.Done()
	s.AcquireN(5)
	for i := 0; i < 1000; i++ {
		sharedCount--
	}
	s.ReleaseN(5)

	fmt.Println("exit decr")
}
