package main

import "runtime"

// Изменение числа ядер на машине
// Изменение привязки приложения к ядрам CPU
// Лимит пропускной способности, основанный на квотах CPU cgroup (linux)

// Sharded data structures by GOMAXPROCS ???

func main() {
	runtime.GOMAXPROCS(10)         // disable container-aware behaviour
	runtime.SetDefaultGOMAXPROCS() // enable container-aware behaviour
}
