package main

import "fmt"

func f1(a string, b int) int {
	fmt.Println("-----f1-----")
	fmt.Println(a, b)
	c := 100
	return c

}

func f2(a string, b int) (int, int) {
	fmt.Println("-----f2-----")
	fmt.Println(a, b)
	return 666, 777
}

func f3(a string, b int) (r1 int, r2 int) {
	fmt.Println("-----f3-----")
	fmt.Println(a, b)
	r1 = 1000
	r2 = 2000
	return
}

func f4(a string, b int) (r1, r2 int) {
	fmt.Println("-----f4-----")
	fmt.Println(a, b)
	r1 = 1000
	r2 = 2000
	return
}

func main() {
	c := f1("abc", 555)
	fmt.Println(c)
	ret1, ret2 := f2("def", 666)
	fmt.Println(ret1, ret2)
	ret3, ret4 := f3("ghi", 777)
	fmt.Println(ret3, ret4)
	ret5, ret6 := f4("jkl", 888)
	fmt.Println(ret5, ret6)
}
