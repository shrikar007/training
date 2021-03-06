package main

import (
	"fmt"
	"program_metrics"
)

func main() {
	metrics.Start()
	c := gen()
	r := calc(c)

	for n := range r {
		fmt.Println(n)
	}

	metrics.End()
}

func gen() chan int {
	c := make(chan int)
	go func() {
		for i:=0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func calc(c chan int) chan int {
	o := make(chan int)
	go func() {
		ic := make(chan int)
		done := make(chan bool)
		var l int
		for n := range c {
			l++
			go func(v int) {
				ic <- v + 1
				done <- true
			}(n)
		}
		go func() {
			for i := 0; i < l; i++ {
				<- done
			}
			close(ic)
		}()
		for n := range ic {
			o <- n
		}
		close(o)
	}()
	return o
}
