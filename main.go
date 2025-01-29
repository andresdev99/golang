package main

type numberTypes interface {
	~int | ~float32
}

func sum[T numberTypes](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	var a myAlias = 42
	sum(a, 5)
}
