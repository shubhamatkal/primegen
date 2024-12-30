# primegen

[![Go Report Card](https://goreportcard.com/badge/github.com/shubhamatkal/primegen)](https://goreportcard.com/report/github.com/shubhamatkal/primegen)

A high-performance prime number generation library for Go that implements both regular and segmented Sieve of Eratosthenes algorithms.

## Features

- Efficient prime number generation up to large numbers
- Automatically switches between regular and segmented sieve based on input size
- Utilizes multiple CPU cores for parallel processing
- Memory efficient implementation for large numbers using segmented sieve

## Installation

```bash
go get github.com/shubhamatkal/primegen
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/shubhamatkal/primegen"
)

func main() {
    // Generate primes up to 100
    primes := primegen.GeneratePrimes(100)
    fmt.Println(primes)
}
```

## Algorithm Details

The library implements two algorithms:

1. **Standard Sieve of Eratosthenes**: Used for numbers up to 10^7
   - Efficient for smaller numbers
   - Uses parallel processing for marking multiples

2. **Segmented Sieve of Eratosthenes**: Used for numbers larger than 10^7
   - Memory efficient for large numbers
   - Processes numbers in segments
   - Utilizes parallel processing for segment calculations

## Performance

The library automatically chooses the most efficient algorithm based on the input size:
- For n ≤ 10^7: Standard Sieve
- For n > 10^7: Segmented Sieve

## Performance Statistics
Below is a comparison of execution times for the standard algorithm versus this optimized implementation in Go:

| Limit           | Standard Segmented | This Go Script Segmented | Non-Segmented Standard  | This Go Script Non-Segmented |
|-----------------|-------------------------|------------------------------|---------------------|--------------------------|
| 10,000,000      | 65.346455 ms            | 38.34653 ms                  | 41.40387 ms         | 37.52741 ms             |
| 100,000,000     | 481.153563 ms           | 438.387912 ms                | 884.653602 ms       | 659.169185 ms           |
| 1,000,000,000   | 4.867800215 s           | 3.215119958 s                | 11.220634894 s      | 10.384302883 s          |
| 9,000,000,000   | 33.484074665 s                    | 29.289063023 s                           | N/A      | N/A          |

## Note

- This comparision is done on AMD Ryzen 5 3450U with 8gb ram
- For 8gb ram Non segmented can handle prime generation upto 8 * 10^9 ; while segmented can handle upto 64 * 10^18
 



## License

MIT License - see LICENSE file

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## References
- [Wikipedia: Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)
