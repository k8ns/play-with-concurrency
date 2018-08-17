package java_util_concurrent


type Runnable interface {
	Run()
}

type Callable interface {
	Call() (error, interface{})
}


type Feature interface {
	Cancel() bool
	IsCancelled() bool
	IsDone() bool
	Get() (error, interface{})
}


type Executor interface {
	Execute(r Runnable)
}

type ExecutorService interface {
	Executor
	Submit(c Callable) (error, Feature)
}
