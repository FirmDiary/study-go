package main

import (
	"fmt"
	"time"
)

//协程
func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)

	ch1 := make(chan string)
	//默认情况下，通信是同步且无缓冲的
	//在有接受者未接收数据之前，发送不会结束
	//所以如果是一个无缓冲的通道，且没有接收数据时，那么通道就会阻塞
	go pump(ch1)
	go suck(ch1) //通过suck去接收了数据
	time.Sleep(1e9)
	//一秒输出了一万多个
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"

	//协程信息标识符
	//表示用通道ch发送变量
}

func getData(ch chan string) {
	var input string

	//time.Sleep(2e9)
	for {
		input = <-ch
		//协程信息标识符
		//表示用input接收通道ch传递的值
		fmt.Printf("%s ", input)
	}
}

func pump(ch chan string) {
	for i := 0; ; i++ {
		ch <- "赵泽武大帅哥"
	}
}

func suck(ch chan string) {
	for {
		fmt.Println(<-ch)
	}
}
