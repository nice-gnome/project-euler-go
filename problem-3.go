/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import "math"

func largestPrimeFactor(n int) int {
	var factors []int
	for n%2 == 0 {
		factors = append(factors, 2)
		n = n / 2
	}
	for i := 3; i <= int(math.Sqrt(float64(n+1))); i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}
	if n > 2 {
		factors = append(factors, n)
	}
	return factors[len(factors)-1]
}

func main() {
	// test
	println(largestPrimeFactor(13195) == 29)

	// solution
	println(largestPrimeFactor(600851475143))
}
