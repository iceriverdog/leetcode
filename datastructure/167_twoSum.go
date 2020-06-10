package datastructure

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers)-1

	for left < right {
		sum:=numbers[left]+numbers[right]
		if target == sum {
			return []int{left+1, right+1}
		} else if  sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}