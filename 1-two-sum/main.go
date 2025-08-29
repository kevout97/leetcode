package main

import "fmt"

func twoSum(nums []int, target int) []int {
	pairs := make(map[int]int)

	for i, num := range nums {
		d := target - num // Get the difference between the current number and the target

		if pairs[d] != 0 {
			return []int{pairs[d] - 1, i}
		}

		pairs[num] = i + 1
	}

	return []int{}
}

func main() {
	fmt.Printf("%v", twoSum([]int{2, 7, 11, 15}, 9))
}
