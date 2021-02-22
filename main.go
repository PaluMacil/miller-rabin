package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	var primes []int
	for i := 500_000; i < 1_000_000; i++ {
		if Prime(int64(i), 10) {
			primes = append(primes, i)
		}
	}

	fmt.Println("primes", primes)
}

func Prime(n int64, rounds int) bool {
	// check low numbers and even numbers...

	if n < 0 {
		panic("n must be positive")
	}
	// 0 and 1 are not prime
	if n < 2 {
		return false
	}
	// 2 and 3 are prime
	if n < 4 {
		return true
	}
	// all even numbers are not prime
	if n&1 == 0 {
		return false
	}
	for i := 0; i < rounds; i++ {
		if MillerRabin(n) {
			return false
		}
	}
	return true
}

// MillerRabin takes a positive integer, n, and returns false for 'inconclusive' or true for 'composite'
func MillerRabin(n int64) bool {
	q := n-1
	k := 0
	for 1 == (q % 2) {
		k += 1
		q = q / 2
	}

	a := rand.Int63() % n
	if 1 == ModularExp(a, q, n) {
		return false // prime
	}
	e := q
	for i := 0; i < k; i++ {
		if (n-1) == ModularExp(a, e, n) {
			return false // prime
		}
		e = 2*e
	}

	return true // composite
}

func ModularExp(base, exponent, n int64) int64 {
	base = base % n
	result := int64(1)
	for i := int64(0); i < exponent; i++ {
		result = (result * base) % n
	}
	return result
}