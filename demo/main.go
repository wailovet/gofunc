package main

import (
	"log"
	"time"

	"github.com/wailovet/gofunc"
)

func main() {
	threadNum := 10
	gofunc.DefaultCatch = func(i interface{}) {

		log.Println("DefaultCatch")
	}
	group := gofunc.NewWaitGroup()
	group.Catch(func(err interface{}) {
		panic(err)
	})
	for ti := 0; ti < threadNum; ti++ {
		func(index int) {
			group.Add(func() {
				time.Sleep(time.Second)
				panic("123")
			}).Catch(func(i interface{}) {
				panic("Catch")
				log.Fatalln("Catch")
			})
		}(ti)
	}
	group.Wait()
	log.Println("OK")
}
