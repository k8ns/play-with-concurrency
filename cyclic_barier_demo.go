package main

import (
	j "play-with-concurrency/java_util_concurrent"
	"fmt"
)


func main() {
	sharedCount = 0

	sb := j.NewSyslicBarier()

	sb.Await()

}


func sbTest(sb *j.CyclicBarrier, name string) {
	sb.Await()
	fmt.Println(name)
}
