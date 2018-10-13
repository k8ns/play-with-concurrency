package main


import (
	j "play-with-concurrency/java_util_concurrent"
	"fmt"
	"time"
	"errors"
)

func demo4() {

	executor := j.GetThreadPool()


	f1, _ := executor.Submit(&CallableOne{"five", 100})
	f2, _ := executor.Submit(&CallableOne{"seven", 250})

	f3, _ := executor.Submit(&CallableTwo{5, 100})
	f4, _ := executor.Submit(&CallableTwo{7, 350})

	r1, _ := f1.Get()
	r2, _ := f2.Get()
	fmt.Println("results one: ", r1, r2)

	r3, _ := f3.Get()
	r4, _ := f4.Get()

	i3 := r3.(int)
	i4 := r4.(int)


	fmt.Println("results two: ", i3 + i4)


	f5, _ := executor.Submit(&CallableErr{})
	r5, r5err := f5.Get()
	fmt.Println("res 5:", "val:", r5, "err:", r5err)
}


type CallableOne struct {
    Value string
    Sleep time.Duration
}


func (c *CallableOne) Call() (interface{}, error) {
	time.Sleep(c.Sleep * time.Millisecond)
	return c.Value, nil
}

type CallableTwo struct {
	Value int
	Sleep time.Duration
}

func (c *CallableTwo) Call() (interface{}, error) {
    time.Sleep(c.Sleep * time.Millisecond)
	return c.Value, nil
}

type CallableErr struct {}

func (c *CallableErr) Call() (interface{}, error) {
	time.Sleep(300 * time.Millisecond)
	return nil, errors.New("some error")
}
