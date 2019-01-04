package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("Holle world init from here...")
}

func main() {
	// syntax.MapTest()
	// syntax.ForTest()
	// syntax.SwitchTest()

	// sum := syntax.ArgsCount(1, 3, 5, 7, 9)
	// fmt.Println("sum =", sum)

	// a := 10
	// fmt.Println("\n", a)
	// syntax.ArgsTest(&a)
	// fmt.Println(a)

	//aStr := "test"
	//fmt.Println("\n", aStr)
	//syntax.ArgsTest2(&aStr)
	//fmt.Println(aStr)

	//read.ReadDup()

	//method.Test()

	//interfaceL.Test()
	//interfaceL.TestFunc()
	// interfaceLearn.TestFunc2()

	//goroutine.Test()

	// web.Test()

	go println("Goroutine test...")
	time.Sleep(time.Millisecond)

}
