package main

func Generator[T any](s []T) <-chan T {
	panic("not implemented")
}

func FanIn[T any](channels []chan<- T) <-chan T {
	panic("not implemented")
}

func FanOut[T any](ch <-chan T) []<-chan T {
	panic("not implemented")
}

// ???
//func WorkerPool()
