package main

import (
	"fmt"
	"runtime"
	"time"
)

// Изменение числа ядер на машине
// Изменение привязки приложения к ядрам CPU
// Лимит пропускной способности, основанный на квотах CPU cgroup (linux)

func main() {
	fmt.Println("Initial GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("NumCPU:", runtime.NumCPU())

	for i := range 100 {
		fmt.Printf("%02d: GOMAXPROCS=%d NumCPU=%d\n", i, runtime.GOMAXPROCS(0), runtime.NumCPU())
		time.Sleep(2 * time.Second)
	}
}
