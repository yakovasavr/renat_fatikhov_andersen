package main

func otrezok (arr, x) (int, int) {
	mapa := make(map[int]int)
	
	i := 0
	sum := 0
	for (i < len(arr)) {
		sum := sum + arr[i]
		mapa[sum] = i
		if mapa[sum - x] {
			fmt.Println(mapa[sum - x] + 1, mapa[sum])
		}
	   i++
	}
	
   fmt.Println(-1, -1)
}

func main() {
	var arr = []int{1,2,3,4,5}
	otrezok(arr, 7)
}


// 0,1,2,34
// 1,2,3,4,5
// 1- >[0]
// 3->[1]
// 6->[2]
// 10->[3]
// -----
// 15[4]
