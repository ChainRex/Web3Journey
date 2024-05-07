package main

import "fmt"

// 定义枚举
const (
	// 每行的iota都会自动递增1
	BEIJING  = 10 * iota // 0
	SHANGHAI             // 10
	SHENZHEN             // 20
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = 1, b = 2
	c, d                      // iota = 1, c = 2, d = 3
	e, f                      // iota = 2, e = 3, f = 4
	g, h = iota * 2, iota * 3 // iota = 3, g = 6, h = 9
	i, k                      // iota = 4, i = 8, k = 12
)

func main() {
	// 常量
	const length int = 10

	fmt.Println(length)

	// length = 100 无法修改

	fmt.Println(BEIJING)
	fmt.Println(SHANGHAI)
	fmt.Println(SHENZHEN)

	// 格式化输出
	fmt.Printf("a = %d,b = %d\n", a, b)
	fmt.Printf("c = %d,d = %d\n", c, d)
	fmt.Printf("e = %d,f = %d\n", e, f)
	fmt.Printf("g = %d,h = %d\n", g, h)
	fmt.Printf("i = %d,k = %d\n", i, k)

	// var a int = iota // iota只能在常量的表达式中使用

}
