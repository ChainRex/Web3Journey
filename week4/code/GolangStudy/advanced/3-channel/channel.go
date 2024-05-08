package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 运行中")
		c <- 666 // 将666发送到channel c
	}()

	num := <-c // 从channel c接收数据，并赋值给num

	fmt.Println("num =", num)

	////////////
	// 有缓存 //
	///////////
	fmt.Println("==========有缓存的channel=============")
	ca := make(chan int, 3) // 带有缓冲的channel
	fmt.Printf("len(ca)=%d, cap(ca)=%d\n", len(ca), cap(ca))
	go func() {
		defer fmt.Println("子goroutine结束")
		for i := 0; i < 4; i++ {
			ca <- i
			fmt.Printf("子goroutine正在运行，发送的元素=%d\n", i)
			fmt.Printf("len(ca)=%d, cap(ca)=%d\n", len(ca), cap(ca))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ca
		fmt.Printf("num=%d\n", num)
	}
	time.Sleep(2 * time.Second)

	//////////////////
	// channel 关闭 //
	/////////////////
	fmt.Println("==========channel关闭=============")

	can := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			can <- i
		}
		close(can) // 关闭channel
		// can <- 666 // 不能再发送数据到channel
	}()
	// // 方法一：for循环遍历channel
	// for {
	// 	if num, ok := <-can; ok { // channel关闭并且channel中的数据都读取完,ok为false
	// 		fmt.Println("num=", num)
	// 	} else {
	// 		break
	// 	}
	// }

	// 方法二：for range遍历channel
	for num := range can {
		fmt.Println("num=", num)
	}
	////////////////////
	// channel select //
	////////////////////
	fmt.Println("==========channel select=============")
	c1 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c1)

		}
		quit <- 0
	}()

	fibonacci(c1, quit)

	fmt.Println("main goroutine 结束")

}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
