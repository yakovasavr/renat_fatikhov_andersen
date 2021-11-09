package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scanf("%d\n", &n)
	if n == 0{
		return
	}
	j := make([]int, n)
	fmt.Scanf("%d\n", &j[0])
	min := j[0]
	minIndex := 0
	i := 1
	for i < n {
		fmt.Scanf("%d\n", &j[i])
		if j[i] < min {
			min = j[i]
			minIndex = i
		}
		i++
	}
	i = minIndex + 1
	sum := 500
	sumprev := 500
	for (i < n - 1) {
		if (j[i] > j[i - 1] || j[i] > j[i + 1]) {
			sum = sum + sumprev + 500
			sumprev = sumprev + 500
		} else if (j[i] == j[i - 1]) {
			sum = sum + 500
		} else if (j[i] < j[i - 1]) {
			sum = sum + 500
		}
		i++
	}
	if j[i] > j[i - 1] {
		sum = sum + sumprev + 500
		sumprev = sumprev + 500
	} else {
		sum = sum + 500
	}
	i = minIndex - 1
	sumprev = 500
	for (i >= 0) {
		if (j[i] > j[i + 1]) {
			sum = sum + sumprev + 500
			sumprev = sumprev + 500
		}
		if (j[i] == j[i + 1]) {
			sum = sum + 500
		}
		i--
	}
	fmt.Printf("%d", sum)
}
