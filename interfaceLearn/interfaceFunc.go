package interfaceLearn

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("Interface init from here...")
}

type Any interface{}

type item struct {
	Name string
}

func (i item) String() string {
	return fmt.Sprintf("item name: %v", i.Name)
}

type person struct {
	Name string
	Sex  string
}

func (p person) String() string {
	return fmt.Sprintf("Person name: %v, sex is: %v", p.Name, p.Sex)
}

func Parse(a Any) Any {
	switch a.(type) {
	case string:
		return &item{Name: a.(string)}
	case []string:
		data := a.([]string)
		if len(data) == 2 {
			return &person{
				Name: data[0],
				Sex:  data[1],
			}
		} else {
			return nil
		}
	default:
		panic(errors.New("type match miss"))
	}
	return nil
}

func TestFunc() {
	p1 := Parse("Apple")
	fmt.Println(p1)
	p2 := Parse([]string{"zhangsan", "man"}).(*person)
	fmt.Println(p2)
}

// 接口型函数
// *********************************旧的方式***********************************
// 不用接口型函数：如果想要实现某个接口，必须先定义一个类型，然后让该类型实现该接口里面的函数（函数名字必须与接口里面的相同）；
type Handler interface {
	Do(k, v Any)
}

type welcome string

func (w welcome) Do(k, v Any) {
	fmt.Printf("%s, 我是%s, 今年%d岁\n", w, k, v)
}

func Each(m map[Any]Any, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

//****************************************************************************
//使用接口型函数：
//首先定义一个函数类型HandlerFunc，该类型实现接口Handler里面的函数Do；
//然后其他相同类型不同名字的函数（selfInfo）都可以该接口；也就是函数实现接口了，称为接口型函数；
//然后可以将新函数（selfInfo）通过HandlerFunc(selfInfo)转换为函数类型HandlerFunc，这样可以作为参数传给其他函数了。

type HandlerFunc func(k, v Any)

func (hf HandlerFunc) Do(k, v Any) {
	hf(k, v)
}

func selfInfo(k, v Any) {
	fmt.Printf("大家好, 我是%s, 今年%d岁\n", k, v)
}

//通用强制类型转换
func EachFunc(f func(k, v Any), m map[Any]Any) {
	Each(m, HandlerFunc(f))
}

//**************************************************************************

func TestFunc2() {
	//方式1：
	fmt.Println("旧的方式：\n类型w实现接口Handler的Do方法：")
	persons := make(map[Any]Any)
	persons["张三"] = 20
	persons["李四"] = 30
	persons["王五"] = 40
	w := welcome("大家好")
	Each(persons, w)

	//方式2：
	fmt.Println("\n函数类型HandlerFunc实现接口的Do方法，相同类型函数selfInfo强转为函数类型HandlerFunc，然后作为参数：")
	Each(persons, HandlerFunc(selfInfo))

	//方式3：
	fmt.Println("\n接口型函数+统一的类型转换实现接口：")
	EachFunc(selfInfo, persons)

}
