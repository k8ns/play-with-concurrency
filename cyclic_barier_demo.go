package main

import (
	j "play-with-concurrency/java_util_concurrent"
	"fmt"
    "time"
)


func main() {

	cb := j.NewCyclicBarier(3, reachBarrier)

	go sbTest(cb, "one", 50)
	go sbTest(cb, "two", 50)
	go sbTest(cb, "three", 50)
    /// ---
	go sbTest(cb, "four", 50)
	go sbTest(cb, "five", 50)
	go sbTest(cb, "six", 50)
    /// ---
    go sbTest(cb, "seven", 50)
    go sbTest(cb, "eight", 50)
    go sbTest(cb, "nine", 50)
	/// ---


	time.Sleep(time.Second)

	fmt.Println(cb.Len())
	fmt.Println("Done")
}


func reachBarrier() {
    fmt.Println("reach barrier")
}

func sbTest(sb *j.CyclicBarrier, name string, w time.Duration) {
    time.Sleep(w * time.Millisecond)
    fmt.Println(name)
    sb.Await()

}
