// https://leetcode.com/problems/missing-number/

package main

import "fmt"

func missingNumber(nums []int) int {
	sum := 0
	sumReal := 0
	i := 0
	for i < len(nums) {
		sum = sum + nums[i]
		sumReal = sumReal + i
		i++
	}
	sumReal = sumReal + i
	return (sumReal - sum)
}

func main() {
	arr := []int{1, 0, 3}
	fmt.Println(missingNumber(arr))
}