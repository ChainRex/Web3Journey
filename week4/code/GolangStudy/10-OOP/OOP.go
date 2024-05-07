package main

import "fmt"

// 声明一个新的类型
type myint int

type Book struct {
	title  string
	author string
}

func changeBook(book Book) {
	// 传递的是副本
	book.author = "mary"
}

func changeBook2(book *Book) {
	// 传递的是指针
	book.author = "mary"
}

type Hero struct {
	name  string
	ad    int
	level int
}

// 方法是作用于特定类型的函数
func (hero *Hero) show() {
	fmt.Printf("Hero name=%v, ad=%v, level=%v\n", hero.name, hero.ad, hero.level)
}

func (hero *Hero) setName(newName string) {

	hero.name = newName
}

//////////
// 继承 //
/////////
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human Eat()")
}

func (this *Human) Walk() {
	fmt.Println("Human Walk()")
}

///////////////////////////////

type SuperMan struct {
	Human // 继承Human的方法和属性
	level int
}

func (this *SuperMan) Eat() {
	fmt.Println("SuperMan Eat()")
}

func (this *SuperMan) Fly() {
	fmt.Println("SuperMan Fly()")
}

func (this *SuperMan) Print() {
	fmt.Println("name = ", this.name)
	fmt.Println("sex = ", this.sex)
	fmt.Println("level = ", this.level)
}

//////////
// 多态 //
/////////

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat Sleep()")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog Sleep()")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

// interface{} 万能数据类型
func myfunc(arg interface{}) {
	fmt.Println("myfunc is called")
	fmt.Println(arg)

	// 断言
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string", value)
	} else {
		fmt.Println("arg is not string")
	}
}

func main() {
	var a myint = 10
	fmt.Printf("a = %v\n", a)
	fmt.Printf("a type is %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.author = "tom"
	fmt.Printf("book1=%v\n", book1)
	changeBook(book1)
	fmt.Printf("book1=%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("book1=%v\n", book1)

	hero := Hero{name: "zhang3", ad: 50, level: 1}

	hero.show()
	hero.setName("li4")
	hero.show()

	h := Human{"zhang3", "female"}
	h.Eat()
	h.Walk()

	s := SuperMan{Human{"zhang3", "female"}, 1}

	s.Walk() // 父类的方法
	s.Eat()  // 子类重写的方法
	s.Fly()  // 子类的方法

	s.Print()

	// 多态
	var animal AnimalIF // 接口指针
	animal = &Cat{"white"}
	animal.Sleep()
	animal = &Dog{"black"}
	animal.Sleep()

	cat := Cat{"yellow"}
	dog := Dog{"black"}
	showAnimal(&cat)
	showAnimal(&dog)

	// interface{} 万能数据类型
	myfunc(100)
	myfunc(cat)
	myfunc("100")
}
