package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 || (n > 2 && n%2 == 0) {
		return false
	}

	for i := 3; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	n, primes := 10, 0
	i := 2

	for primes < n {
		if isPrime(i) {
			fmt.Printf("%d ", i)
			primes++
		}
		i++
	}
}
