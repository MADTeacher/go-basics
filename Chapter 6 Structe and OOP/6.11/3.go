package main

import "fmt"

type MyInt int

func Sum[T int | float64 | string](a, b T) T {
	return a + b
}

func main() {
	var value1, value2 MyInt = 34, 22
	fmt.Println(Sum(value1, value2))
	// MyInt does not implement int|float64|string
	// (possibly missing ~ for int in constraint int|float64|string)
}
