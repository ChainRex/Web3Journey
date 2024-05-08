package main

import "fmt"

func main() {
	// 写入defer关键字，类似析构函数,在函数return之后执行
	defer fmt.Println("Goodbye, World!1")
	defer fmt.Println("Goodbye, World!2") // 先执行
	fmt.Println("Hello, World!")
}
