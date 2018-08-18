package java_util_concurrent

import "errors"

var (
	NullPointerErr  = errors.New("nil")
	RejectedExecutionErr  = errors.New("rejected")
)

type Runnable interface {
	Run()
}

type Callable interface {
	Call() (interface{}, error)
}




type Executor interface {
	Execute(r Runnable)
}

type ExecutorService interface {
	Executor
	Submit(c Callable) (Feature, error)
}

func GetThreadPool() ExecutorService {
	return &ExecutorServiceImplementation{}
}



type ExecutorServiceImplementation struct {

}

func (s *ExecutorServiceImplementation) Execute(r Runnable) {

}

func (s *ExecutorServiceImplementation) Submit(c Callable) (Feature, error) {

	if c == nil {
		return nil, NullPointerErr
	}

	f := &FeatureImplementation{false, nil, make(chan interface{}), make(chan error)}

	go func(){
		v, err := c.Call()
		if err != nil {
			f.errCh <- err
		} else {
			f.ch <- v
		}
	}()

	return f, nil
}
