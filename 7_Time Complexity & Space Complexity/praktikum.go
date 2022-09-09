package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(pow(2, 3))
	// fmt.Println(pow(5, 3))
	// fmt.Println(pow(10, 2))
	// fmt.Println(pow(2, 5))
	// fmt.Println(pow(7, 3))

	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))

}

func primeNumber(number int) bool {

	if number == 1 {
		return false

	} else if number % 2 == 0 {
		return false
	}
	
	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
			if number % i == 0 {
				return false
		}
	}

	return true
}

func pow(x, n int) int {
	var nilai int
	var total int = 1
	nilai = x
	for i := 0; i < n; i++ {
		total = total * nilai
	}
	return total
}
