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
