package java_util_concurrent


type Feature interface {
	//Cancel() bool
	//IsCancelled() bool
	IsDone() bool
	Get() (interface{}, error)
}

type FeatureImplementation struct {
	done bool
	val interface{}
	ch chan interface{}
}

func (f *FeatureImplementation) IsDone() bool {
	return f.done
}

func (f *FeatureImplementation) Get() (interface{}, error) {

	if f.val != nil {
		return f.val, nil
	}

	cycle: for {
		select {
			case f.val = <-f.ch:
				f.done = true
				close(f.ch)
				break cycle
		}
	}

	return f.val, nil
}
