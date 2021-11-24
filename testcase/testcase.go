// https://github.com/KazanExpress/golang-test-task

package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// ErrBlocked reports if service is blocked.
var ErrBlocked = errors.New("blocked")

// Service defines external service that can process batches of items.
type Service interface {
	GetLimits() (n uint64, p time.Duration)
	Process(ctx context.Context, batch Batch) error
}

// Batch is a batch of items.
type Batch []Item

// Item is some abstract item.
type Item struct{}

func proccessing(s Service, ctx context.Context, batch Batch){
	err := s.Process(ctx, batch)
	if err == nil{
		fmt.Println(err)
	}
}

// a structure for limitation inputs
type inputValues struct{
	n	uint64		// number of items before block
	d	time.Duration		// duration of proccessing
}

// variable which realizes an interface
type serviceStaging struct{}

func main() {
	var (
		x 		serviceStaging				// variable which realizes an interface
		input 	inputValues = inputValues{}	// initializing new structure for limitation input
		setup 	Batch = Batch{}				// batch - slice of elements
		ctx 	context.Context
	)

	// stage1 getting the limits
	input.n, input.d = x.GetLimits()	//putting values into structure
	itemchannel := make(chan Item, input.n)		//initialization of buffered channel

	// stage2 gathering the Items and processing
	var i uint64 = 0 // counter for fullfilling a Batch
	start := time.Now()	// fixing the start point of time
	for element := range itemchannel{		// element contains 1 Item, received from buffered channel
		setup = append(setup, element)		// appenfing new value into slice of Items
		i++
		// flushing the Batch
		if (i == input.n){
			t := time.Now()
			elapsed := t.Sub(start)
			if (elapsed < input.d){			// counting the time difference
				time.Sleep(input.d - elapsed)		// sleeping the goroutine till we get the input.d parameter	
			}				
			proccessing(x, ctx, setup)	// processing of Items
			i = 0
			setup = setup[:0]			// clearing out a slice
			start = time.Now()			// resetting a time counter
		}
	}
}