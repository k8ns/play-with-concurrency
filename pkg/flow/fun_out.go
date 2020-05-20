package flow

func funOut(In <-chan int, outA, outB, outC chan int) {
	for data := range In {
		select {
		case outA <- data:
		case outB <- data:
		case outC <- data:
		}
	}
	//close(outA)
	//close(outB)
	//close(outC)
}
