package exec

import (
	"sync"
)

// Trigger f1 and f2 in separate goroutines and wait for completition using sync.WaitGroup.
func Await2Tasks[R1 any, R2 any](f1 func() R1, f2 func() R2) (r1 R1, r2 R2) {
	var wg sync.WaitGroup

	wg.Go(func() {
		r1 = f1()
	})

	wg.Go(func() {
		r2 = f2()
	})

	wg.Wait()

	return r1, r2
}

// RunWorker ranges over workCh until channel get's closed and sends work result into resultCh. It doesn't trigger new worker goroutine by default, so run it as 'go func RunWorker(workCh, resultCh)' in case you need one.
func RunWorker[R any](workCh <- chan func() R, resultCh chan <- R) {
	for w := range workCh {
		resultCh <- w()
	}
}