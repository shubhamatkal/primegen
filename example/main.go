package main

import (
	"fmt"
	"time"

	"github.com/shubhamatkal/primegen"
)

func main() {
	// Example 1: Generate small primes
	fmt.Println("Example 1: Small primes (up to 20)")
	smallPrimes := primegen.GeneratePrimes(20)
	fmt.Printf("Primes up to 20: %v\n", smallPrimes)
	fmt.Printf("Count: %d\n\n", len(smallPrimes))

	// Example 2: Medium range primes with timing
	fmt.Println("Example 2: Medium range primes (up to 10000)")
	start := time.Now()
	mediumPrimes := primegen.GeneratePrimes(10000)
	duration := time.Since(start)
	fmt.Printf("Found %d primes up to 10000\n", len(mediumPrimes))
	fmt.Printf("First 10 primes: %v\n", mediumPrimes[:10])
	fmt.Printf("Last 10 primes: %v\n", mediumPrimes[len(mediumPrimes)-10:])
	fmt.Printf("Time taken: %v\n\n", duration)

	// Example 3: Large range using segmented sieve
	fmt.Println("Example 3: Large range primes (up to 1000000)")
	start = time.Now()
	largePrimes := primegen.GeneratePrimes(1000000)
	duration = time.Since(start)
	fmt.Printf("Found %d primes up to 1000000\n", len(largePrimes))
	fmt.Printf("First 10 primes: %v\n", largePrimes[:10])
	fmt.Printf("Last 10 primes: %v\n", largePrimes[len(largePrimes)-10:])
	fmt.Printf("Time taken: %v\n\n", duration)

}
