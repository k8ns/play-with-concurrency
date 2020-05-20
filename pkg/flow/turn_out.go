package flow

import "fmt"

func turnOutfunnel(InA, InB, InC <-chan int, outA, outB, outC chan int) {
	go funnel(InA, InB, InC, outA)
	go funnel(InA, InB, InC, outB)
	go funnel(InA, InB, InC, outC)

	fmt.Println("exit from turn out (Funnel version)")
}

func turnOutOut(InA, InB, InC <-chan int, outA, outB, outC chan int) {
	go funOut(InA, outA, outB, outC)
	go funOut(InB, outA, outB, outC)
	go funOut(InC, outA, outB, outC)

	fmt.Println("exit from turn out (Fun Out version)")
}

func turnOut(InA, InB, InC <-chan int, outA, outB, outC chan int) {

	var data int
	var open bool
	var openA bool
	var openB bool
	var openC bool

	for {
		select {
		case data, openA = <-InA:
			open = openA
		case data, openB = <-InB:
			open = openB
		case data, openC = <-InC:
			open = openC
		default:
			fmt.Println("read def")
			continue
		}

		if !openA && !openB && !openC {
			fmt.Println("all closed. Break")
			break
		}

		if !open { // no data
			fmt.Println("no data")
			continue
		}

		select {
		case outA <- data:
		case outB <- data:
		case outC <- data:
		}
	}

	fmt.Println("exit from turn out")

	close(outA)
	close(outB)
	close(outC)

}
