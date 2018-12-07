package interfaceL

import (
	"fmt"
)

//定义Interface
type Men interface {
	SayHi()
	Sing(lyrics string)
}

type YongChap interface {
	SayHi()
	Sing(lyrics string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(lyrics string)
	SpendSalary(amount float32)
}

//定义结构体
type Human struct {
	name  string
	age   int8
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s, you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Printf("La la, la la la, la la la la la, %s...\n", lyrics)
}

func (s Student) SayHi() {
	fmt.Printf("Hi, I am %s, I am a Student, you can call me on %s\n", s.name, s.phone)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
	fmt.Printf("Student's amount: %f...\n", s.loan)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I am an Employee, you can call me on %s\n", e.name, e.phone)
}

func interfaceValue() {
	mark := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}

	//定义Men类型的变量m
	var m Men
	//m存储Student
	m = mark
	m.SayHi()
	m.Sing("test")

	value, ok := m.(Student)
	fmt.Println("对象值: ", value, "对象是否是Student：", ok)

	//m存储Employee
	m = tom
	m.SayHi()
	m.Sing("test")

	value, ok = m.(Student)
	fmt.Println("对象值: ", value, "对象是否是Student：", ok)

	//定义slice men
	menSlice := [2]Men{}
	menSlice[0], menSlice[1] = mark, tom

	for _, value := range menSlice {
		value.SayHi()
	}
}

func Test() {
	h := Human{"Jack", 28, "123456"}
	fmt.Println("")
	h.SayHi()
	h.Sing("wo shi mai bao de xiao hang jia")
	fmt.Println("")

	s := Student{Human{"Tom", int8(22), "1234567788"}, "qinghua", 0}
	s.SayHi()
	s.BorrowMoney(100)
	fmt.Println("")

	interfaceValue()
}

// 总结：从上面可以看到Men interface被Human、Student 和Employee实现。
// 同理,一个对象可以实现任意多个interface,例如上面的Student实现了Men和YonggChap两个interface。
