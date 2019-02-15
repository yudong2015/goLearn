package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/structs"
	s "github.com/yudong2015/goLearn/syntax"
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

	fmt.Println("Test envs...")
	fmt.Println(os.Getenv("PATH"))
	//go println("Goroutine test...")
	time.Sleep(time.Millisecond)


	var f4 float64 = 1.2345678
	var f5 float64 = f4 * 100 /100


	fmt.Println(f4)
	fmt.Println(f5)
	fmt.Printf("%.2f\n", f5)
	fmt.Printf("%.3f\n", f5)
	fmt.Printf("%.4f\n", f5)

	fmt.Println(structs.Name(&s.TestStruct{}))
}
