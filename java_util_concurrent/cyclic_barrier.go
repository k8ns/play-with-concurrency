package java_util_concurrent

import (
    "sync"
    "time"
)

type CyclicBarrier struct {
	numOfRoutines int32
	action func()
	count int32
	chans []chan bool
	lock sync.Locker
    trackLock sync.Locker
}



func NewCyclicBarier(numOfRoutines int32, action func()) *CyclicBarrier {
    return &CyclicBarrier{
        numOfRoutines,
        action,
        0,
        make([]chan bool, 0),
        &sync.Mutex{},
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
    //fmt.Println(len(cb.chans))
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
    cb.trackLock.Lock()
    cb.count++
    release := cb.count == 3
    if release {
       cb.count = 0
    }

    //atomic.AddInt32(&cb.count, 1)
    //release := atomic.CompareAndSwapInt32(&cb.count, cb.numOfRoutines, 0)
    if release {
        cb.action()
        chRelease := cb.chPop()
        for _, ch := range chRelease {
            ch <- true
            close(ch)
        }

        time.Sleep(10 * time.Millisecond)
    }


    cb.trackLock.Unlock()
}

