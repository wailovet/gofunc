package gofunc

import (
	"sync"
	"time"
)

var DefaultCatch = func(interface{}) {

}

type goFun struct {
	exception func(interface{})
}

func (gf *goFun) Catch(handle func(interface{})) *goFun {
	gf.exception = handle
	return gf
}

type goFuncLoop struct {
	gf     *goFun
	isQuit bool
}

func (gfl *goFuncLoop) Quit() {
	gfl.isQuit = true
}

func (gfl *goFuncLoop) Catch(handle func(interface{})) *goFuncLoop {
	gfl.gf.exception = handle
	return gfl
}

func Loop(f func(i uint64) bool) *goFuncLoop {
	gfl := goFuncLoop{
		gf: &goFun{},
	}
	go func() {
		defer func() {
			// defer func() { recover() }()

			if e := recover(); e != nil {
				if gfl.gf.exception != nil {
					gfl.gf.exception(e)
				} else {
					DefaultCatch(e)
				}
			}
		}()

		var i uint64
		for i = 0; !gfl.isQuit && f(i); i++ {

		}
	}()
	return &gfl
}

func New(f func()) *goFun {
	gfl := Loop(func(i uint64) bool {
		f()
		return false
	})
	return gfl.gf
}

type process struct {
	isWait bool
}

func (p *process) pause() {
	p.isWait = true
	for p.isWait {
		time.Sleep(time.Second / 2)
	}
}

func (p *process) Play() {
	p.isWait = true
}

func Pause() *process {
	p := process{}
	p.pause()
	return &p
}

type waitGroup struct {
	swg             sync.WaitGroup
	processTotal    int
	processCount    int
	onProcessHandle func(i float32)
	exception       func(interface{})
}

func NewWaitGroup() *waitGroup {
	return &waitGroup{}
}

func (wg *waitGroup) Add(f func()) *waitGroup {
	wg.processTotal++
	wg.swg.Add(1)
	New(func() {
		defer func() {
			wg.swg.Done()
			wg.processCount++
			if wg.onProcessHandle != nil {
				wg.onProcessHandle((float32(wg.processCount) / float32(wg.processTotal)) * 100)
			}
		}()
		f()
	}).Catch(func(i interface{}) {
		if wg.exception != nil {
			wg.exception(i)
		}
	})
	return wg
}

func (wg *waitGroup) Catch(handle func(interface{})) *waitGroup {
	wg.exception = handle
	return wg
}

func (wg *waitGroup) OnProcess(handle func(i float32)) *waitGroup {
	wg.onProcessHandle = handle
	return wg
}

func (wg *waitGroup) Wait() {
	wg.swg.Wait()
}
