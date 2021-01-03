package main

import (
	"fmt"
	"time"
)

func fib(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
			return 1
		} else {
			return fib(n-1) + fib(n-2)
		}
}

func main() {
	start := time.Now()

	n := uint(10)
	fmt.Println(fib(n))

	elapsed := time.Since(start)
	fmt.Println("Binomial took %s", elapsed)
}
