package testing01

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums{
		sum += n
	}
	return sum
}