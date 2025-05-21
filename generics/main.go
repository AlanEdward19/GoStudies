package main

func main() {
	result := add(5, 10)
	println(result) // Output: 15
}

func add[T int | float64](a, b T) T {
	return a + b
}
