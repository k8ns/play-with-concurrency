package java_util_concurrent

import (
	"sync/atomic"
	"time"
	"sync"
)

type Semaphore struct {
	acquiredCount int32
	lock *sync.Mutex
}

func NewSemaphore() *Semaphore {
	return &Semaphore{
		0,
		&sync.Mutex{},
	}
}

func (s *Semaphore) Acquire() {
	s.AcquireN(1)
}

func (s *Semaphore) AcquireN(n int32) {
	for {
		if atomic.LoadInt32(&s.acquiredCount) == 0 {
			return
		}
	}

	atomic.AddInt32(&s.acquiredCount, n)
	time.Sleep(100 * time.Millisecond)
}

func (s *Semaphore) Release() {
	s.ReleaseN(s.acquiredCount)
}

func (s *Semaphore) ReleaseN(n int32) {
	atomic.AddInt32(&s.acquiredCount, -n)
}
