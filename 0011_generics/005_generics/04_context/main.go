package main

// go build main.go

func main() {
	m := make(map[any]any)
	insert(m, "@igoroutine")
	insert(m, 1)
	insert(m, 2)
	insert(m, struct{}{})

	println(len(m))

}

func insert[T any](m map[any]any, a T) {
	type Data struct{}
	d := Data{}

	m[d] = a
}
