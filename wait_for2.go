package main

import (
	"fmt"
	"runtime"
	"sync"
)

func print(wg *sync.WaitGroup, n int) {
	x := 0
	for i := 0; i < 100; i++ {
		x += i
	}
	fmt.Println(n, x)
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go print(&wg, i)
	}

	wg.Wait()
	fmt.Println("over")

}
