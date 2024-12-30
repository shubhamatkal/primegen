// Package primegen provides efficient prime number generation using both standard
// and segmented Sieve of Eratosthenes algorithms.
//
// The package automatically chooses between two implementations based on the input size:
// - Standard Sieve of Eratosthenes for numbers up to 10^7
// - Segmented Sieve of Eratosthenes for larger numbers
//
// Both implementations utilize parallel processing for better performance on
// multi-core systems.
package primegen
