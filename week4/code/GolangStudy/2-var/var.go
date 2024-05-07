package main

import "fmt"

// 声明全局变量

var gA = 100
var gB int = 200

// gC := 300 // 全局变量声明不能使用简短声明

func main() {
	// 默认值是0
	var a int
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b int = 100
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s,type of bb = %T\n", bb, bb)

	// 自动匹配类型
	var c = 100
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	var cc = "abcd"
	fmt.Printf("cc = %s,type of cc = %T\n", cc, cc)

	// 简短声明
	d := 100
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	f := "abcd"
	fmt.Printf("f = %s,type of f = %T\n", f, f)

	g := 3.14
	fmt.Printf("g = %f,type of g = %T\n", g, g)

	// 全局变量
	fmt.Println(gA)
	fmt.Println(gB)

	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println(xx, yy)
	var kk, ll = 100, "abcd"
	fmt.Println(kk, ll)

	// 多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println(vv, jj)
}
