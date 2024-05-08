package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {

	fmt.Println("Type:", reflect.TypeOf(arg))
	fmt.Println("Value:", reflect.ValueOf(arg))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("User Call")
	fmt.Printf("%+v\n", *this)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input).Elem()
	fmt.Println("inputType is:", inputType.Name())
	// 获取input的value
	inputValue := reflect.ValueOf(input).Elem()
	fmt.Println("inputValue is:", inputValue)

	// 分别通过type获取里面的字段
	// 1.通过type获取NumField,进行遍历
	// 2.得到每个field,数据类型
	// 3.通过field的interface(),获取对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type获取里面的方法，获取指针类型
	inputType = reflect.TypeOf(input)
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}

}

///////////////
// 结构体标签 //
///////////////

type resume struct {
	Name string `info:"name" doc:"my name"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	fmt.Println("t is:", t)
	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		fmt.Printf("taginfo is:%s\n", taginfo)
		docinfo := t.Field(i).Tag.Get("doc")
		fmt.Printf("docinfo is:%s\n", docinfo)
	}
}

//////////////
// json应用 //
/////////////

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actor"`
}

func main() {
	var num float64 = 3.4
	reflectNum(num)

	user := User{1, "Midori", 12}
	user.Call()
	DoFiledAndMethod(&user)
	///////////////
	// 结构体标签 //
	///////////////

	re := resume{"Midori", "M"}

	findTag(&re)

	//////////////
	// json应用 //
	/////////////

	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "张柏芝"}}

	// 结构体转json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error")
		return
	}
	fmt.Printf("jsonStr is:%s\n", jsonStr)

	// json转结构体

	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error")
		return
	}

	fmt.Printf("myMovie is:%+v\n", myMovie)

}
