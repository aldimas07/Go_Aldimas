package main

import (
	"fmt"
	"math"

	// "sort"
)

func main() {
	// searching()

	// fmt.Println(primeX(1))
	// fmt.Println(primeX(5))
	// fmt.Println(primeX(8))
	// fmt.Println(primeX(9))
	// fmt.Println(primeX(10))

	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(12))
}

func primeX(number int) int {
	if number == 1 {
		return 2
	} else if number < 1 {
		return 0
	}

	var jumlah int = 1
	var i int = 1

	for jumlah <= number{
		if cekBilPrima(i){
			jumlah++
		}
		i++
	}
	return i-1

}


func cekBilPrima(number int) bool {
	var bagi int
	if number == 1 {
		return false
	} else if number%2 == 0 {
		return false
	}

	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {

		if number%i == 0 {
			bagi = bagi + 1
		}
	}

	if bagi == 2 {
		return false
	} else {
		return true
	}

}
func fibonacci(number int) int {
	
	if number <= 1 {
		return number
	}

	return fibonacci(number-1) + fibonacci(number-2)

}

// func searching() bool {
// 	element := []int{11, 22, 55, 33, 44, 12, 54, 66}
// 	nilai := 66

// 	pencarian := sort.SearchInts(element, nilai)

// 	if nilai == element[pencarian] {
// 		res := true
// 		fmt.Println("Ketemu", element[pencarian], res)
// 		return res
// 	} else {
// 		fmt.Println("Tidak ketemu", element[pencarian])
// 		res := false
// 		return res
// 	}
// }
