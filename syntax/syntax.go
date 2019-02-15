package syntax

import "fmt"

// Array : test go array
func Array() {
	fmt.Println("About array: ")

	//var ar [10]int8
	ar := [10]int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var aSlice, bSlice []int8

	aSlice = ar[3:5]
	fmt.Println(aSlice)
	fmt.Println(len(aSlice))

	aSlice = ar[:]
	fmt.Println(aSlice)

	aSlice = ar[3:7]
	fmt.Println(aSlice)

	bSlice = aSlice[1:3]
	fmt.Println(bSlice)

	bSlice = aSlice[3:]
	fmt.Println(bSlice)
	fmt.Println(len(bSlice))
	fmt.Println(cap(bSlice))

	fmt.Println("")

	bSlice = append(bSlice, int8(10))
	fmt.Println(ar)
	fmt.Println(aSlice)
	fmt.Print("len: ")
	fmt.Println(len(bSlice))
	fmt.Print("cap: ")
	fmt.Println(cap(bSlice))
	fmt.Println(bSlice)
	fmt.Println(bSlice[len(bSlice)-1])

}

func MapTest() {
	fmt.Println("\nAbout map: ")
	//声明map-1
	var aNums map[string]int
	aNums = make(map[string]int)

	//声明map-2
	bNums := make(map[string]int)

	aNums["a"] = 10
	bNums["a"] = 100

	fmt.Println(aNums)
	fmt.Println(aNums["a"])
	fmt.Println(bNums)
	fmt.Println(bNums["a"])

	aNums["m"] = 30
	fmt.Println(aNums)
	delete(aNums, "m")
	fmt.Println(aNums)

}

func ForTest() {
	fmt.Println("\nAbout for: ")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	mm := make(map[string]string)
	mm["a"] = "aaaa"
	mm["b"] = "bbbb"

	for k, v := range mm {
		fmt.Println(k + ": " + v)
	}

}

func SwitchTest() {
	fmt.Println("\nAbout switch: ")
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1.")
	case 2:
		fmt.Println("i is equal to 2.")
	case 9, 10, 11:
		fmt.Println("i is in 9, 10, 11.")
	default:
		fmt.Println("switch default...")
	}
}

func ArgsCount(args ...int) int {
	fmt.Println("\nAbout variable args: ")
	sum := 0
	for _, i := range args {
		sum += i
	}
	return sum
}

func ArgsTest(a *int) int {
	fmt.Println("About pre args1: ")
	*a = *a + 1
	return *a
}

func ArgsTest2(a *string) string {
	fmt.Println("About pre args2: ")
	*a = *a + "111"
	return *a
}

type TestStruct struct {
	name  string
}