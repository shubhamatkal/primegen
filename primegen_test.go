package primegen

import (
	"testing"
)

// Known prime numbers for different ranges
var (
	primesUnder10  = []uint64{2, 3, 5, 7}
	primesUnder20  = []uint64{2, 3, 5, 7, 11, 13, 17, 19}
	primesUnder100 = []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
)

func TestGeneratePrimes(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected []uint64
	}{
		{"Empty for 1", 1, []uint64{}},
		{"Only 2 for input 2", 2, []uint64{2}},
		{"Primes under 10", 10, primesUnder10},
		{"Primes under 20", 20, primesUnder20},
		{"Primes under 100", 100, primesUnder100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GeneratePrimes(tt.input)
			if !equalSlices(got, tt.expected) {
				t.Errorf("GeneratePrimes(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

// Helper functions
func equalSlices(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
