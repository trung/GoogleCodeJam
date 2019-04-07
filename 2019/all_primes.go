package main

import "fmt"

func sieve(n int) {
	primes := make([]bool, n+1)
	for i := range primes {
		primes[i] = true
	}
	for p := 2; p*p < n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}
	for i := 2; i <= n; i++ {
		if primes[i] {
			fmt.Print(i, ",")
		}
	}
}

func main() {
	sieve(10000)
}
