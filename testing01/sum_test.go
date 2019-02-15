package testing01

import (
	"testing"
	"fmt"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := 45

	actual := Sum(numbers)

	if actual != expected {
		t.Errorf("The sum of %v is %d, expected is %d.", numbers, actual, expected)
	}
}


//example test
func ExampleSum() {
	numbers := []int{1,3,5,7}
	fmt.Println(Sum(numbers))
	//Output:
	//16
}