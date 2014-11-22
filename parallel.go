package main

import "fmt"

type Vector []float64

func (v Vector) DoSome(i int, c chan int) {
	fmt.Println("doSome", i)
	c <- 1 // 发送完成信号
}

//我们在一个循环中为每个CPU启动一个独立的计算片段，这些片段可以以任意的顺序执行，执行顺序在这里是无关紧要的。
//在启动所有的goroutines后，我们只需要从信道中提取所有的完成信号即可。

const NCPU = 4 // CPU核数

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // Buffering optional but sensible.
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i, c)
	}
	//从信道中取出所有信号
	for i := 0; i < NCPU; i++ {
		<-c // 等待一个任务完成
	}
	fmt.Println("finish all")
}

func main() {
	var v Vector = Vector{10, 20, 30}
	v.DoAll(v)
}
