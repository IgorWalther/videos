package main

import (
	"time"
)

type Data struct {
	First  int
	Second int
}

func main() {
	for range 100 {
		var data Data
		values := make([]int, 2)

		//wg := new(sync.WaitGroup)

		// [G1 G1 G1 G1]...[G2, G2, G2, G2]
		go func() {
			data.First = 27
			values[0] = 10
		}()

		go func() {
			data.Second = 28
			values[1] = 11
		}()

		//wg.Go(func() {
		//	data.First = 27
		//	values[0] = 10
		//})
		//
		//wg.Go(func() {
		//	data.Second = 28
		//	values[1] = 11
		//})

		//wg.Wait()

		time.Sleep(time.Millisecond * 200)

		if data.First != 27 || data.Second != 28 || values[0] != 10 || values[1] != 11 {
			panic("wtf")
		}

		//
		//fmt.Println(data.First)
		//fmt.Println(data.Second)
		//
		//fmt.Println(values[0])
		//fmt.Println(values[1])

		// Will we see it here?
	}
}
