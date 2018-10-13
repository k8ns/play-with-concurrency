package java_util_concurrent

import (
    "sync"
)

type CyclicBarrier struct {
	numOfRoutines int32
	action        func()

    lock          sync.Locker
	chans         []chan bool

    count         int32
	trackLock     sync.Locker
}

func NewCyclicBarrier(numOfRoutines int32, action func()) *CyclicBarrier {
	return &CyclicBarrier{
        numOfRoutines: numOfRoutines,
		action: action,
		count: 0,
		chans: make([]chan bool, 0),
		lock: &sync.Mutex{},
        trackLock: &sync.Mutex{},
	}
}

func (cb *CyclicBarrier) Len() int {
	return len(cb.chans)
}

func (cb *CyclicBarrier) Await() {
	ch := cb.chCreate()
	go cb.track()
	<-ch
}

func (cb *CyclicBarrier) chCreate() chan bool {
	ch := make(chan bool)
	cb.lock.Lock()
	cb.chans = append(cb.chans, ch)
	cb.lock.Unlock()
	return ch
}

func (cb *CyclicBarrier) track() {
	cb.trackLock.Lock()
	defer cb.trackLock.Unlock()

	cb.count++
	release := cb.count == cb.numOfRoutines
	if !release {
	    return
    }

	cb.count = 0

    chRelease := cb.chPop()
    for _, ch := range chRelease {
        ch <- true
        close(ch)
    }

    cb.action()
}

func (cb *CyclicBarrier) chPop() []chan bool {
	cb.lock.Lock()

	defer cb.lock.Unlock()

	ret := cb.chans[:cb.numOfRoutines]
	cb.chans = cb.chans[cb.numOfRoutines:]
	return ret
}
