package java_util_concurrent



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



	f := &FeatureImplementation{false, nil, make(chan interface{})}

	go func(){
		v, _ := c.Call()
		f.ch <- v
	}()

	return f, nil
}
