package main

import "fmt"

func main2() {
	// var primes [10]int

	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// primes[3] = 7
	// primes[4] = 11
	// primes[5] = 13
	// primes[6] = 17
	// primes[7] = 19
	// primes[8] = 23
	// primes[9] = 29

	// Array mit fest vorgegebener Länge
	// primes := [8]int{2, 3, 5, 7, 11, 13, 17, 19}

	// Array mit fester Länge, die inferiert wird
	// primes := [...]int{2, 3, 5, 7, 11, 13, 17, 19}

	// Slice mit variabler Länge
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	// primes := make([]int, 0, 1000)
	// fmt.Println(len(primes), cap(primes))

	primes = append(primes, 23)
	primes = append(primes, 29)

	// for i := 0; i < len(primes); i++ {
	// 	fmt.Println(primes[i])
	// }

	for _, prime := range primes {
		fmt.Println(prime)
	}

	population := map[string]int{}
	population["Berlin"] = 3_769_495
	population["Hamburg"] = 1_852_478
	population["München"] = 1_450_381

	for city, pop := range population {
		fmt.Printf("%s: %d\n", city, pop)
	}
}
