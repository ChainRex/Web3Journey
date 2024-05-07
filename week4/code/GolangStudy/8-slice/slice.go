package main

import "fmt"

// 函数传递定长数组，要声明数组长度,值拷贝
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 引用传递
func printArray2(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
	arr[0] = 100
}

func main() {

	// 固定长度
	var array1 [5]int
	array2 := [10]int{1, 2, 3, 4, 5} // 后面没有赋值的默认为0
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	for index, value := range array2 {
		fmt.Printf("array2[%d] = %d\n", index, value)
	}

	// 查看类型
	fmt.Printf("%T\n", array1)
	fmt.Printf("%T\n", array2)
	printArray(array1)

	array3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 动态数组，切片
	fmt.Printf("%T\n", array3)
	printArray2(array3) // 引用传递
	for _, v := range array3 {
		fmt.Println(v)
	}

	var array4 []int
	fmt.Println(len(array4))
	array4 = make([]int, 3) // 开辟空间
	fmt.Println(len(array4))

	var array5 []int = make([]int, 3)
	fmt.Println(len(array5))

	array6 := make([]int, 3)
	fmt.Println(len(array6))

	// 判断一个切片是否为空
	var array7 []int
	if array7 == nil {
		fmt.Println("array4 is nil")
	} else {
		fmt.Println("array4 is not nil")
	}

	var array8 = make([]int, 3, 5) // 长度为3，容量为5（默认为len）,append容量空间不够时会开辟cap数量的容量
	fmt.Printf("len(array8) = %d,cap(array8) = %d,slice = %v\n", len(array8), cap(array8), array8)
	array8[1] = 2
	array8[2] = 3
	array8 = append(array8, 1)
	fmt.Printf("len(array8) = %d,cap(array8) = %d,slice = %v\n", len(array8), cap(array8), array8)
	slice := array8[1:3] // 切片截取，左闭右开，引用同一个数组，修改切片会影响原数组
	fmt.Printf("len(slice) = %d,cap(slice) = %d,slice = %v\n", len(slice), cap(slice), slice)

	slice2 := make([]int, 3)
	copy(slice2, array8[1:3]) // 拷贝
	fmt.Printf("len(slice2) = %d,cap(slice2) = %d,slice2 = %v\n", len(slice2), cap(slice2), slice2)
}
