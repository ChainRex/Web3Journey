package main

import "fmt"

// 引用传递
func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	var m1 map[string]string // map[key]value
	if m1 == nil {
		println("m1 is nil")
	}
	// 使用map前，需要先make，分配数据空间
	m1 = make(map[string]string, 10)
	m1["name"] = "tom"
	m1["age"] = "20"

	fmt.Printf("m1=%v\n", m1)

	m2 := make(map[int]string)
	m2[1] = "tom"
	m2[2] = "jack"

	fmt.Printf("m2=%v\n", m2)

	m3 := map[string]string{
		"name": "tom",
		"age":  "20",
	}
	fmt.Printf("m3=%v\n", m3)

	// 遍历
	for k, v := range m3 {
		fmt.Println(k, v)
	}
	// 删除
	delete(m3, "name")
	fmt.Printf("m3=%v\n", m3)

	// 修改
	m3["age"] = "30"
	fmt.Printf("m3=%v\n", m3)

	printMap(m3)
}
