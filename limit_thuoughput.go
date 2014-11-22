package main

import "fmt"
import "time"

type Request struct {
	name string
	id   int64
}

var MaxOutstanding = 2
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
	sem <- 1   // 等待队列缓冲区非满
	process(r) // 处理请求，可能会耗费较长时间.
	<-sem      // 请求处理完成，准备处理下一个请求
}

func Serve(queue chan *Request) {
	for {
		req := <-queue
		go handle(req) //不等待handle完成，不阻塞
	}
}

func process(r *Request) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(r.id, r.name)
}

func main() {
	var queue = make(chan *Request, 1)

	//for i := 0; i < 10; i++ {
	//produce msg
	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			now := time.Now()
			i := now.Format("2006-01-02 15:04:05") //time.Now().Unix()
			r := Request{
				fmt.Sprint("name: ", i),
				now.Unix(),
			}

			fmt.Println("start push queue..", i)
			queue <- &r
			fmt.Println("finish push queue..", i)

		}
	}()

	Serve(queue)
}
