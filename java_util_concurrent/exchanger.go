package java_util_concurrent

import (
	"sync"
)

type Exchanger struct {
	party int
	ch1 chan string
	ch2 chan string
	lock sync.Locker
}

func NewExchanger() *Exchanger {
	return &Exchanger{
		0,
		make(chan string),
		make(chan string),
		&sync.Mutex{},
	}
}

func (e *Exchanger) Exchange(v string) string {

	n := 0
	e.lock.Lock()
	e.party++
	if e.party == 1 {
		n = 1
	}
	if e.party == 2 {
		e.party = 0
		n = 2
	}
	e.lock.Unlock()

	if n == 1 {
		return e.one(v)
	}

	if n == 2 {
		return e.two(v)
	}

	return ""
}

func (e *Exchanger) one(v string) string {
	go func(){
		e.ch1 <- v
		close(e.ch1)
	}()
	return <-e.ch2
}

func (e *Exchanger) two(v string) string {
	go func() {
		e.ch2 <- v
		close(e.ch2)
	}()
	return <-e.ch1
}
