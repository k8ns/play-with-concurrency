package java_util_concurrent

import (
    "sync/atomic"
    "sync"
)

type CyclicBarrier struct {
	numOfRoutines int32
	action func()
	count int32
	chans []chan bool
	lock sync.Locker
}



func NewCyclicBarier(numOfRoutines int32, action func()) *CyclicBarrier {
    return &CyclicBarrier{
        numOfRoutines,
        action,
        0,
        make([]chan bool, 0),
        &sync.Mutex{},
    }
}

func (cb *CyclicBarrier) Len() int {
    return len(cb.chans)
}

func (cb *CyclicBarrier) chCreate() chan bool {
    ch := make(chan bool)
    cb.lock.Lock()
    cb.chans = append(cb.chans, ch)
    cb.lock.Unlock()
    return ch
}

func (cb *CyclicBarrier) chPop() []chan bool {
    cb.lock.Lock()
    defer cb.lock.Unlock()
    ret := cb.chans[:cb.numOfRoutines]
    cb.chans = cb.chans[cb.numOfRoutines:]
    return ret
}

func (cb *CyclicBarrier) Await() {
    ch := cb.chCreate()
    go cb.track()
    <-ch
}

func (cb *CyclicBarrier) track() {
    atomic.AddInt32(&cb.count, 1)
    if atomic.CompareAndSwapInt32(&cb.count, cb.numOfRoutines, 0) {
        go cb.action()
        chRelease := cb.chPop()
        for _, ch := range chRelease {
            ch <- true
            close(ch)
        }
    }
}

