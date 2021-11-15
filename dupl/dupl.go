// Find the Duplicate Number
// https://leetcode.com/problems/find-the-duplicate-number/

package main

import "fmt"

type tree struct {
    val int
    left *tree
    right *tree
}

func add(t *tree, num int, d *int) *tree {
	if t == nil {
        t = new(tree)
		   t.val = num
        return t
    }
	if num == t.val {
		*d = num
	}
    if num < t.val {
        t.left = add(t.left, num, d)
    }
    if num > t.val {
        t.right = add(t.right, num, d)
    }
	return t
}

func findDuplicate(nums []int) int {
	var t *tree
	var d int
    for _, num := range nums {
		t = add(t, num, &d)
    }
	return d
}

func main(){
	arr := []int{4, 4, 4, 4}
	fmt.Println(findDuplicate(arr))
}
