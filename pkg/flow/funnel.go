package flow

func funnel(InA, InB, InC <-chan int, out chan int) {
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
		}

		if !openA && !openB && !openC {
			break
		}

		if !open { // no data
			continue
		}

		out <- data
	}

	close(out)
}
