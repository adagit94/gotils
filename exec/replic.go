package exec

import (
	"sync"
)

func Replicate[A any, V any](maxRetries uint8, op func(arg A) (V, error), args ...A) ([]V, []error) {
	values := make([]V, len(args))
	errors := make([]error, len(args))

	var wg sync.WaitGroup

	for i, arg := range args {
		wg.Go(func() {
			val, err := Retried(func() (V, error) {
				return op(arg)
			}, maxRetries)

			values[i] = val
			errors[i] = err
		})
	}

	wg.Wait()

	return values, errors
}
