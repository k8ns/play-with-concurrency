package java_util_concurrent

import (
    "sync/atomic"
)


type CountDownLatch struct{
    count int32
    release chan bool
    countReached bool
}

func NewCountDownLatch(count int32) *CountDownLatch{
    return &CountDownLatch{
        count,
        make(chan bool),
        false,
    }
}

func (c *CountDownLatch) CountDown() {
    atomic.AddInt32(&c.count, -1)
    if atomic.LoadInt32(&c.count) == 0 {
        c.release <- true
        close(c.release)
    }
}

func (c *CountDownLatch) Await() {
    <-c.release
}
