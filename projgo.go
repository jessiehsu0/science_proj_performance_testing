package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	//fmt.Println(sum)
     
        elapsed := time.Since(start)
	fmt.Println("Binomial took %s", elapsed)
}
