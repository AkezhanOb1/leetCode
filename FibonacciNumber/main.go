package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func fib(N int) int {
	previous := 0
	current := 1
	if N <= 1 {
		return N
	}

	for i := 0; i < N-1; i++ {
		current = previous + current
		previous = current - previous
	}

	return current
}
