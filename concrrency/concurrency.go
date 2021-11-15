package main

// import (
// 	"fmt"
// )

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int, 2)

	go func() {
		ch1 <- 4
	}()

	go func() {
		v := <- ch1
		ch2 <- v + 1
		ch2 <- 4
		ch2 <- 6
	}()

	// fmt.Println(<-ch2)
	// fmt.Println(<-ch2)
	// fmt.Println(<-ch2)


}
