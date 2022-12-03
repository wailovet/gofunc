package gofunc

import "github.com/gammazero/workerpool"

type Pool struct {
	wp     *workerpool.WorkerPool
	doFunc func(in interface{})
}

func NewPool(t int, doFunc func(in interface{})) (pool *Pool) {
	pool = &Pool{}
	pool.wp = workerpool.New(t)
	pool.doFunc = doFunc
	return pool
}

func (pool *Pool) Do(in interface{}) {
	pool.wp.Submit(func() {
		pool.doFunc(in)
	})
}

func (pool *Pool) Wait() {
	pool.wp.StopWait()
}
