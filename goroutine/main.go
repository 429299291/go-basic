package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(i int) {
	for a := 0; a < 10; a++ {
		fmt.Println("test()", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器➖1
}

//<<<<<<<<<<<<<<<<<<<<<<<<<<goroutine和 channel结合，一个写数据一个读取数据，同步
var wg2 sync.WaitGroup

func fn1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("写入数据", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
	wg2.Done()
}
func fn2(ch chan int) {
	for v := range ch {
		fmt.Println("读取数据", v)
		time.Sleep(time.Millisecond * 10)
	}
	wg2.Done()
}
func main() {
	for a := 0; a < 10; a++ {
		wg.Add(1)  //协程计数➕1
		go test(a) //开启协程
	}
	wg.Wait() //等待协程执行完毕
	fmt.Println("主线程执行完毕")
	//<<<<<<<<<<<<<<<<<<<<<<<<<< channel	引用数据类型,指针地址， 先入先出的原则	|	先关闭管道 close(ch1)    for range 遍历，没有 key
	ch1 := make(chan int, 3)
	ch1 <- 4
	c1 := <-ch1
	fmt.Println(ch1, cap(ch1), len(ch1), c1)
	//<<<<<<<<<<<<<<<<<<<<<<<<<<goroutine和 channel结合，一个写数据一个读取数据，同步
	ch := make(chan int, 10)
	wg2.Add(2)
	go fn1(ch)
	go fn2(ch)
	wg2.Wait()
}
