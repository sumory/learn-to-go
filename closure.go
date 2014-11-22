package main

import "fmt"

func adder() func(int, int) (int, int) {
	sum := 0
	i := 0
	return func(x, y int) (int, int) {
		sum += x + y
		i++
		return sum, i
	}
}

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a

	}
}

func main() {
	pos := adder()
	for i := 0; i < 3; i++ {
		x, y := pos(i, i)
		fmt.Println("value:", x, "\tn:", y)
	}

	println("fibonacci func...")
	//fibonacci
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
