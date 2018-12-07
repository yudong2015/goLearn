package method

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// more complex

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

func Test() {
	fmt.Println("Method simple test...\n ")

	r1 := Rectangle{10, 20}
	r2 := Rectangle{9.9, 3.5}
	c1 := Circle{10}
	c2 := Circle{3.3}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())

}
