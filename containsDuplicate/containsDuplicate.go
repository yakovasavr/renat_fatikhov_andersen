package main

import "fmt"

func containsDuplicate(nums []int) bool {
	for i, value := range nums {
		j := i + 1
		for j < len(nums) {
			if value == nums[j] {
				return true
			}
			j++
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4, 2}
	fmt.Println(containsDuplicate(arr))
}

