package java_util_concurrent

import (
	"fmt"
	"time"
)

func ExampleExecutorAndFeature() {
	executor := GetThreadPool()

	f4, _ := executor.Submit(&CallableObj{"four", 400})
	f2, _ := executor.Submit(&CallableObj{"two", 200})
	f5, _ := executor.Submit(&CallableObj{"five", 500})
	f3, _ := executor.Submit(&CallableObj{"three", 300})
	f1, _ := executor.Submit(&CallableObj{"one", 100})

	r1, _ := f1.Get()
	r2, _ := f2.Get()
	r3, _ := f3.Get()
	r4, _ := f4.Get()
	r5, _ := f5.Get()

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(r5)

	//Output:
	//one
	//two
	//three
	//four
	//five


}

type CallableObj struct {
	Value string
	Sleep time.Duration
}


func (c *CallableObj) Call() (interface{}, error) {
	time.Sleep(c.Sleep * time.Millisecond)
	return c.Value, nil
}