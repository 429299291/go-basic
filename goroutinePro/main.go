package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 0; i < 300000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for v := range intChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- v
		}
	}
	// close(primeChan)	多个协程里如果关闭了，无法继续给 channel 继续写入数据
	exitChan <- true
	wg.Done()
}
func printChan(primeChan chan int) {
	// for v := range primeChan {
	// 	fmt.Println(v)
	// }
	wg.Done()
}

// 查找质数
func main() {
	startTime := time.Now().UnixNano() / 1e6
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 80000) //这个要注意容量大小
	exitChan := make(chan bool, 16)    //用来标识关闭 primeChan 的
	wg.Add(1)
	//存放数字的协程
	go putNum(intChan)
	//统计质数的协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}
	//打印质数的协程
	wg.Add(1)
	go printChan(primeChan)
	//
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()
	wg.Wait()
	endTime := time.Now().UnixNano() / 1e6
	fmt.Println("执行完毕", "用时", endTime-startTime)
}
