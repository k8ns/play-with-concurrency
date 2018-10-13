package main

import (
	j "play-with-concurrency/java_util_concurrent"
	"sync"
	"fmt"
	"time"
)

func demo3() {

	exch := j.NewExchanger()
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func(ex *j.Exchanger) {
		time.Sleep(100 * time.Millisecond)
		s := exch.Exchange("one")
		fmt.Println("one got", s)
		wg.Done()
	}(exch)

	go func(ex *j.Exchanger) {
		time.Sleep(100 * time.Millisecond)
		s := exch.Exchange("two")
		fmt.Println("two got", s)
		wg.Done()
	}(exch)




	wg.Wait()


}
