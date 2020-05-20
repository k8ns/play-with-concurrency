package java_util_concurrent

import (
	"fmt"
	"sync"
)

func ExampleCountDownLatch() {
	latch := NewCountDownLatch(5)

	var wg sync.WaitGroup
	wg.Add(10)

	//second group
	for i := 0; i < 5; i++ {
		go func() {
			latch.Await()
			fmt.Println("second")
			wg.Done()
		}()
	}

	// first group
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("first")
			latch.CountDown()
			wg.Done()
		}()
	}

	wg.Wait()

	// Output:
	// first
	// first
	// first
	// first
	// first
	// second
	// second
	// second
	// second
	// second
}
