package flow

import (
	"time"
	"fmt"
)

func main() {

	InA := make(chan int)
	InB := make(chan int)
	InC := make(chan int)

	outA := make(chan int)
	outB := make(chan int)
	outC := make(chan int)

	a := make([]int, 0)
	b := make([]int, 0)
	c := make([]int, 0)

	go chanListen(outA, &a)
	go chanListen(outB, &b)
	go chanListen(outC, &c)

	//go feedIn(InA, 0, 30)
	//go funOut(InA, outA, outB, outC)

	go feedIn(InA, 0, 10)
	go feedIn(InB, 10, 20)
	go feedIn(InC, 20, 30)
	go turnOut(InA, InB, InC, outA, outB, outC)
	//go turnOutfunnel(InA, InB, InC, outA, outB, outC)
	//go turnOutOut(InA, InB, InC, outA, outB, outC)


	//go funnel(InA, InB, InC, outA)




	time.Sleep(100 * time.Millisecond)



	//close(outA)
	//close(outB)
	//close(outC)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("DONE")

}

func feedIn(In chan int, start, count int) {
	for i := start; i < count; i++ {
		In <- i
	}
	close(In)
}







func chanListen(c <-chan int, a *[]int) *[]int {
	for data := range c {
		*a = append(*a, data)
	}
	return a
}