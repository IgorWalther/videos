package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := range 100 {
		wg.Go(func() {
			fmt.Println(i)
		})
	}

	wg.Wait()
}

//func main() {
//	wg := new(sync.WaitGroup)
//
//	for i := range 100 {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			fmt.Println(i)
//		}()
//	}
//
//	wg.Wait()
//}
