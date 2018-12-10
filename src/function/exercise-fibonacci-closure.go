package main

import "fmt"

func fibonacci() func() int {
	b := -1
	a := 1
	return func() int {
		// b, a = a, a+b
		a, b = a+b, a
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
