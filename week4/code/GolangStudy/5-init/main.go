package main

import (
	_ "GolangStudy/5-init/lib1" // 匿名，无法使用当前包，但是会调用init方法
	// "GolangStudy/5-init/lib2"
	mylib2 "GolangStudy/5-init/lib2"
	// . "GolangStudy/5-init/lib2"
)

func main() {
	// lib1.Lib1Test()
	mylib2.Lib2Test()

}
