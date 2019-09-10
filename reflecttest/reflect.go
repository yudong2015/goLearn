package reflecttest

import (
	"fmt"
	"reflect"
)

type Man interface {
	Say()
}

type User struct {
	Name    string
	age     int
	address string
}

func


func Test(){
	u := User{
		"tom",
		27,
		"wuhan",
	}

	//reflect.ValueOf(u): 返回对象u的值, 如果是指针则返回指针地址
	//reflect.TypeOf(u): 返回对象u的类型，int/string/StructName/...
	//reflect.TypeOf(u).Kind(): 返回对象u类型的类型，int/string/struct/...
	fmt.Println("")
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	fmt.Println("Type: ")
	fmt.Println(t)
	fmt.Println("Kind: ")
	fmt.Println(t.Kind())  //kind is the type of Type ::::::::: struct
	fmt.Println("Value: ")
	fmt.Println(v)
	fmt.Println("")


	v1 := reflect.ValueOf(u.Name)
	t1 := reflect.TypeOf(u.Name)
	fmt.Println("Type: ")
	fmt.Println(t1)
	fmt.Println("Kind: ")
	fmt.Println(t1.Kind())  //kind is the type of Type :::::: string
	fmt.Println("Value: ")
	fmt.Println(v1)
	fmt.Println("")



	//reflect.ValueOf(v).Elem()   只能取指针指向的值 或者 是接口的实现值，如果reflect.Valueof(v)返回值不是指针或者接口则会报错
	//reflect.Indirect(reflect.Valueof(v))   取指针指向的值或者自身，不管reflect.Valueof(v)返回的值是不是指针都可以取到值（不是指针，返回本身）
	i := 20
	fmt.Println("")
	fmt.Println(reflect.TypeOf(i))
	fmt.Println("Before i = ", i)

	vi1 := reflect.ValueOf(&i)
	fmt.Println("vi = ", vi1)
	fmt.Println("e1 = ", vi1.Elem())
	e2 := reflect.Indirect(vi1)
	fmt.Println("e2 = ", e2)
	fmt.Println("")

	vi2 := reflect.ValueOf(i)
	fmt.Println("v = ", vi2)
	//fmt.Println("e1 = ", vi2.Elem())
	e4 := reflect.Indirect(vi2)
	fmt.Println("e2 = ", e4)
	fmt.Println("")
}
