package main

import "fmt"

const (
	A = 0
	C = 0
	D = 0
	B = iota
)

func sliceChange(sl []int) {
	// sl = append(sl, 5)
	sl[0] = 4
}

func sliceChangePointer(sl *[]int) {
	(*sl)[0] = 4
	*sl = append(*sl, 5)
}

func main() {
	fmt.Println(A, C, D, B)
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	sliceChange(a)
	fmt.Println(a, cap(a), &a[0])
	sliceChangePointer(&a)
	fmt.Println(a, cap(a), &a[0])
}