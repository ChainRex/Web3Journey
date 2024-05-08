package main

import (
	"fmt"
	"runtime"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Println("newTask i:", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine,主goroutine退出后所有子goroutine也会退出
func main() {
	// 创建一个goroutine去执行newTask函数
	go newTask()
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 退出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	go func(a int, b int) bool {
		fmt.Printf("a=%d,b=%d\n", a, b)
		return true
	}(10, 20)

	i := 0
	for {
		i++
		fmt.Println("main i:", i)
		time.Sleep(5 * time.Second)
		if i == 5 {
			break
		}
	}
}
