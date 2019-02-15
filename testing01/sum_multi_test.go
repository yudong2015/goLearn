package testing01

import (
	"testing"
)


//multi test data in one func
func TestSumMul(t *testing.T) {
	numbers1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected1 := 45
	numbers2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, -100}
	expected2 := -55

	t.Run("numbers1", testSumMulFunc(numbers1, expected1))
	t.Run("numbers2", testSumMulFunc(numbers2, expected2))

}

func testSumMulFunc(nums []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := Sum(nums)

		if actual != expected {
			t.Errorf("The sum of %v is %d, expected is %d.", nums, actual, expected)
		}
	}
}