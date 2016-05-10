// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// GenPrimesTrialDiv generates a set of prime numbers, using
// trial division method, up to and including number n.
// The nil, false is returned when primes cannot be generated up to the n.
// The time complexity is O(n**3/2). The O(1) additional space
// is needed (beyond the space needed to write the final result).
func GenPrimesTrialDiv(n uint) (primes []uint, ok bool) {
	switch {
	case n < 2:
		return nil, true
	case n*n < n: // Check if primes generation up to the n will overflow.
		return nil, false
	}

	// isPrime walk through already generated primes
	// and check if n is a product of those primes.
	isPrime := func(n uint) bool {
		for _, p := range primes {
			switch {
			case p*p > n: // Check primes up to the sqrt of n.
				return true
			case n%p == 0:
				return false
			}
		}
		return true
	}

	primes = append(primes, 2)
	for i := uint(3); i <= n; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes, true
}

// GenPrimesSieve generates a set of prime numbers, using
// trial division method, up to and including number n.
// The nil, false is returned when primes cannot be generated up to the n.
// The time complexity is O(n*log(log(n))), and O(n) additional space is needed.
func GenPrimesSieve(n uint) (primes []uint, ok bool) {
	var size uint = (n-3)/2 + 1 // For n > 2: n-3: subtract num: 0,1,2; div by 2: we need only to go through odd numbers; +1: include num: n

	switch {
	case n < 2:
		return nil, true
	case n == 2: // In order to avoid using math.Floor function when computing size we handle the one error case here.
		return []uint{2}, true
	case 2*size*size+6*size+3 < size: // Check if primes generation up to the n will overflow.
		return nil, false
	}

	primes = append(primes, 2)
	isNotPrime := make([]bool, size) // Represents whether 2*i + 3 is prime or not. By default all are primes.
	for i := uint(0); i < size; i++ {
		if !isNotPrime[i] {
			p := 2*i + 3
			primes = append(primes, p)

			// Remove all multiplies of p from the candidate set isNotPrime.
			// Sieving from p*p, whose value is (4*i*i + 12*i + 9) because all
			// numbers of the form k*p where k < p have already been sieved out.
			// The index of this value in isNotPrime is (2*i*i + 6i + 3) because isNotPrime[i] represents 2*i + 3.
			j := 2*i*i + 6*i + 3
			for j < size {
				isNotPrime[j] = true
				jn := j + p
				if jn < j { // Check for j+p overflows.
					return nil, false
				}
				j = jn
			}
		}
	}
	return primes, true
}
