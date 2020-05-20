package java_util_concurrent

import (
	"fmt"
	"sync"
	"time"
)

func ExampleCyclicBarrier() {

	var wg sync.WaitGroup
	wg.Add(12)

	cb := NewCyclicBarrier(3, func() {
		time.Sleep(time.Millisecond * 3)
		fmt.Println("barrier")
		wg.Done()
	})

	cbTestHelper := func(cb *CyclicBarrier) {
		cb.Await()
		fmt.Println("action")
		wg.Done()
	}

	for i := 0; i < 9; i++ {
		go cbTestHelper(cb)
	}

	wg.Wait()

	//Output:
	//action
	//action
	//action
	//barrier
	//action
	//action
	//action
	//barrier
	//action
	//action
	//action
	//barrier
}
