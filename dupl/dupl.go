// Find the Duplicate Number
// https://leetcode.com/problems/find-the-duplicate-number/

package main

import "fmt"

type node struct {
    val int
    left *tree
    right *tree
}

type tree struct {
	root *node
	double int
}

// func (t *tree) findDup(num int) int{
//     if t == nil {
//         t.val = num
//        t.left = new(tree)
//        t.right = new(tree)
//         return 0
//     }
//     if t.val == num {
//         return 1
//     }
//     if num < t.val {
//         t.left.findDup(num)
//     }
//     if num > t.val {
//         t.right.findDup(num)
//     }
// 	return 0
// }

// func add(t *tree, num int) *tree {
// 	if t == nil {
//         t = new(tree)
// 		t.val = num
//         return t
//     }
//     if num < t.val {
//         t.left = add(t.left, num)
//     }
//     if num > t.val {
//         t.right = add(t.right, num)
//     }
// 	return t
// }

// func findDup(t *tree, num int) int {
// 	if t == nil {
//         t = new(tree)
// 		t.val = num
// 		return 0
//     } else if num < t.val {
//         findDup(t.left, num)
//     } else if num > t.val {
//         findDup(t.right, num)
//     }
// 	return 2
// }

// func add2(t *tree, num int) int {
// 	if t == nil {
//         t = new(node)
// 		t.val = num
//         return 0
//     }
//     if num < t.val {
//         add2(t.left, num)
//     }
//     if num > t.val {
//         add2(t.right, num)
//     }
// 	return 0
// }

func (n *node) insert(num int, t *tree) {
	if num < n.val {
		if n.left == nil {
			n.left = &node{val: num}
		} else {
			n.left.insert(num, t)
		}
	} else if num > n.val {
		if n.right == nil {
			n.right = &node{val: num}
		} else {
			n.right.insert(num, t)
		}
	} else {
		t.double = num
	}
}

func (t *tree) createRoot(num int) {
	if t.root == nil {
		t.root = &node{val: num}
	} else {
		t.root.insert(num, t.root)
	}
}

func findDuplicate(nums []int) int {
	var t *tree
    for _, num := range nums {
		t.createRoot(num)
        //  if t.double == num {
        //     return num
        // }
    }
	return t.val
}

func main(){
	arr := []int{4, 2, 3, 1}
	fmt.Println(findDuplicate(arr))

}