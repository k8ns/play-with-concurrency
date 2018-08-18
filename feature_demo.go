package main


import (
	j "play-with-concurrency/java_util_concurrent"
	"fmt"
	"time"
)

func main() {

	executor := j.GetThreadPool()


	f1, _ := executor.Submit(&CallableOne{"5", 100})
	f2, _ := executor.Submit(&CallableOne{"7", 1500})


	r1, _ := f1.Get()
	r2, _ := f2.Get()



	fmt.Println("results: ", r1, r2)



}


type CallableOne struct {
	Value string
	Sleep time.Duration
}


func (c *CallableOne) Call() (interface{}, error) {

	time.Sleep(200 * time.Millisecond)

	return c.Value, nil
}