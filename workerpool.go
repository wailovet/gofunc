package gofunc

import "github.com/gammazero/workerpool"

type Pool struct {
	wp *workerpool.WorkerPool
}

func NewPool(t int) (pool *Pool) {
	pool = &Pool{}
	pool.wp = workerpool.New(t)
	return pool
}

func (pool *Pool) Do(doFunc func(in interface{}), in interface{}) {
	pool.wp.Submit(func() {
		doFunc(in)
	})
}

func (pool *Pool) Wait() {
	pool.wp.StopWait()
	pool.wp = workerpool.New(pool.wp.Size())
}
