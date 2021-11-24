//https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	lg := len(nums) + 1
    v := make([]int, lg)
	i := 0
	for i < len(nums) {
		v[nums[i]] = nums[i]
		i++
	}
	i = 1
	k := 0
	for i < lg {
		if v[i] == 0 {
			v[k] = i
			k++
		}
		i++
	}
	return v[:k]
}

func main(){
	arr := []int{4,3,2,7,3,2,3,1}
	fmt.Println(findDisappearedNumbers(arr))
}