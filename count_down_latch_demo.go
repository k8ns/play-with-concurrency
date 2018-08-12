package main

import (
	"sync"
	"strconv"
	"fmt"
	j "play-with-concurrency/java_util_concurrent"
)

func main() {

	cdl := j.NewCountDownLatch(5)

	var wg sync.WaitGroup

	wg.Add(20)

	for i := 0; i < 10; i++ {
		go waiter(cdl, "T" + strconv.Itoa(i), &wg)
	}

	for i := 0; i < 10; i++ {
		go runner(cdl, &wg)
	}

	wg.Wait()

	fmt.Println("Done")
}


func waiter(cl *j.CountDownLatch, Name string, wg *sync.WaitGroup) {
	fmt.Println("waiter", Name, "waits")
	cl.Await()
	fmt.Println("waiter", Name, "run")
	wg.Done()
}

func runner(cl *j.CountDownLatch, wg *sync.WaitGroup) {
	fmt.Println("runner run")
	cl.CountDown()
	wg.Done()
}