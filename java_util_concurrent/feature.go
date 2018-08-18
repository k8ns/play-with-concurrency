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
	errCh chan error
}

func (f *FeatureImplementation) IsDone() bool {
	return f.done
}

func (f *FeatureImplementation) Get() (interface{}, error) {

	if f.val != nil {
		return f.val, nil
	}

	var err error

	cycle: for {
		select {
			case f.val = <-f.ch:
				f.doneit()
				break cycle

			case err = <-f.errCh:
				f.doneit()
				break cycle
		}
	}

	return f.val, err
}


func (f *FeatureImplementation) doneit() {
	f.done = true
	close(f.ch)
	close(f.errCh)
}