package interfaceL

import (
	"errors"
	"fmt"
)

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
type Handler interface {
	Do(k, v Any)
}

func Each(m map[Any]Any, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

type welcome string

func (w welcome) Do(k, v Any) {
	fmt.Printf("%s, 我是%s, 今年%d岁\n", w, k, v)
}

//****************************************************************************

type HandlerFunc func(k, v Any)

func (hf HandlerFunc) Do(k, v Any) {
	hf(k, v)
}

func (w welcome) selfInfo(k, v Any) {
	fmt.Printf("%s, 我是%s, 今年%d岁\n", w, k, v)
}

//**************************************************************************

func EachFunc(f func(k, v Any), m map[Any]Any) {
	Each(m, HandlerFunc(f))
}

func TestFunc2() {
	//方式1：
	fmt.Println("旧的实现接口：")
	persons := make(map[Any]Any)
	persons["张三"] = 20
	persons["李四"] = 30
	persons["王五"] = 40
	w := welcome("大家好")
	Each(persons, w)

	//方式2：
	fmt.Println("\n接口型函数+类型转换实现接口：")
	Each(persons, HandlerFunc(w.selfInfo))

	//方式3：
	fmt.Println("\n接口型函数+统一的类型转换实现接口：")
	EachFunc(w.selfInfo, persons)

}
