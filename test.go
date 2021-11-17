package main

import (
	"fmt"
	"strings"
)

func fibonacci(n uint64) uint64 {
	var (
		f0, f1 uint64 = 0, 1
	)
	if n == 0 {
		return 0
	}
	for i := 1; uint64(i) < n; i++ {
		f0, f1 = f1, f0+f1
	}
	return f1
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d - %d\n", i, fibonacci(uint64(i)))
	}
	str := "asdqweeeww"
	fmt.Print(strings.Count(str, "e"))
}
