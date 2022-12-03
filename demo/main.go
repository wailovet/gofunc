package main

import (
	"log"
	"time"

	"github.com/wailovet/gofunc"
)

func main() {
	pool := gofunc.NewPool(2)
	for i := 0; i < 20; i++ {

		pool.Do(func(in interface{}) {

			log.Println(in)
		}, i)
	}

	pool.Wait()

	log.Println("--------------------------------------------")
	time.Sleep(time.Second)

	for i := 30; i < 50; i++ {
		pool.Do(func(in interface{}) {

			log.Println(in)
		}, i)
	}

	pool.Wait()
}
