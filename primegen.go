package primegen

import (
	"math"
	"runtime"
	"sync"
)

func GeneratePrimes(n uint64) []uint64 {
	// Choose appropriate method based on input size
	var primes []uint64
	if n <= 10000000 { // 10^7
		primes = SieveOfEratosthenes(n)
	} else {
		primes = Segmented_SOE(n)
	}
	return primes
}

// This is without segmenting.
func SieveOfEratosthenes(n uint64) []uint64 {
	cores := runtime.NumCPU()
	next := make(chan bool, cores)
	var nums = make([]bool, n/2+1)
	m := uint64(math.Sqrt(float64(n)))

	for i := uint64(3); i <= m; i = i + 2 {
		if nums[i/2] == false {
			go goFill(nums, i, n, next)
			next <- true
		}
	}

	for i := 0; i < cores; i++ {
		next <- true
	}

	var ps []uint64
	if n >= 2 {
		ps = append(ps, 2)
	}
	for i := uint64(3); i <= n; i = i + 2 {
		if nums[i/2] == false {
			ps = append(ps, i)
		}
	}
	return ps
}

func fill(nums []bool, i uint64, max uint64) {
	// a := 3 * i
	iteration := 0
	a := i * i
	for a <= max {
		iteration++
		nums[a/2] = true
		a = a + 2*i
	}
}

func goFill(nums []bool, i uint64, max uint64, next chan bool) {
	fill(nums, i, max)
	<-next
}

// Segmented Sieve
var csegPool sync.Pool

func fillSegments(n uint64, basePrimes []uint64, allPrimes *[]uint64, segSize uint64, segNum uint64, next chan bool, nextTurn []chan bool) {
	cseg := (csegPool.Get()).([]bool)
	for i := uint64(0); i < segSize; i++ {
		cseg[i] = false
	}

	segEnd := segSize * (segNum + 1)

	for i := 0; i < len(basePrimes); i++ {
		p := basePrimes[i]
		pSquare := p * p

		if pSquare > segEnd {
			continue
		}

		jMax := segSize * (segNum + 1) / basePrimes[i]

		startJ := basePrimes[i] - 1
		if startJ < (segSize*segNum)/basePrimes[i] {
			startJ = (segSize * segNum) / basePrimes[i]
		}

		for j := startJ; j < jMax; j++ {
			sn := (j + 1) * basePrimes[i]
			cseg[sn-segSize*segNum-1] = true
		}
	}

	if segNum > 1 {
		<-nextTurn[segNum]
	}

	for i := uint64(0); i < segSize; i++ {
		if !cseg[i] && segSize*segNum+i+1 <= n {
			*allPrimes = append(*allPrimes, segSize*segNum+i+1)
		}
	}

	<-next
	if int(segNum)+1 < len(nextTurn) {
		nextTurn[segNum+1] <- true
	}

	csegPool.Put(cseg)
}

func Segmented_SOE(n uint64) (allPrimes []uint64) {
	allPrimes = make([]uint64, 0, n/uint64(math.Log(float64(n))-1))

	segSize := uint64(math.Sqrt(float64(n)))

	csegPool.New = func() interface{} {
		return make([]bool, segSize)
	}

	basePrimes := SieveOfEratosthenes(segSize)
	allPrimes = append(allPrimes, basePrimes...)

	cores := runtime.NumCPU()
	next := make(chan bool, cores)
	var nextTurn []chan bool
	nextTurn = make([]chan bool, n/segSize+1)
	for i := uint64(0); i < n/segSize+1; i++ {
		nextTurn[i] = make(chan bool)
	}
	for segNum := uint64(1); segNum <= n/segSize-1; segNum++ {
		go fillSegments(n, basePrimes, &allPrimes, segSize, segNum, next, nextTurn)
		next <- true
	}
	for i := 0; i < cores; i++ {
		next <- true
	}

	return allPrimes
}
