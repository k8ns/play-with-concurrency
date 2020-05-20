package java_util_concurrent

import (
	"fmt"
	"sync"
)

func ExampleStringExchanger() {

	exch := NewExchanger()
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		s := exch.Exchange("one")
		fmt.Println("one got", s)
		wg.Done()
	}()

	go func() {
		s := exch.Exchange("two")
		fmt.Println("two got", s)
		wg.Done()
	}()

	wg.Wait()

	//Unordered output:
	//one got two
	//two got one
}
