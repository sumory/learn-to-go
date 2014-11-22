package main

import (
	"fmt"
	"runtime"
)

func print(c chan bool, n int) {
	x := 0
	for i := 0; i < 100; i++ {
		x += i
	}
	fmt.Println(x)
	c <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go print(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}

	fmt.Println("over")
}
