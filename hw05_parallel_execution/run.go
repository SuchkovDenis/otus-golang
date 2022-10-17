package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}

	var wg sync.WaitGroup
	var errLim int32
	ch := make(chan Task, len(tasks))
	for _, task := range tasks {
		ch <- task
	}
	close(ch)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range ch {
				if atomic.LoadInt32(&errLim) >= int32(m) {
					return
				}
				if task() != nil {
					atomic.AddInt32(&errLim, 1)
				}
			}
		}()
	}

	wg.Wait()

	if errLim >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
