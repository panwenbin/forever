package forever

import "sync"

// Go runs your function in several routines, with a mutex
func Go(run func(mux *sync.Mutex), routineLimits int) {
	mux := &sync.Mutex{}
	for i := 0; i < routineLimits; i++ {
		go func() {
			for {
				run(mux)
			}
		}()
	}
}
