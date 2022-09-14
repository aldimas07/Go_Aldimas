package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// searching()

	// fmt.Println(primeX(1))
	// fmt.Println(primeX(5))
	// fmt.Println(primeX(8))
	// fmt.Println(primeX(9))
	// fmt.Println(primeX(10))

	// fmt.Println(fibonacci(0))
	// fmt.Println(fibonacci(2))
	// fmt.Println(fibonacci(9))
	// fmt.Println(fibonacci(10))
	// fmt.Println(fibonacci(12))

	// primaSegiEmpat(2, 3, 13)
	// primaSegiEmpat(5, 2, 1)

	// fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	// fmt.Println(MaxSequence([]int{-2, 5, 6, -2, -3, 1, 5, -6}))
	// fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	// fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	// fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))

	// fmt.Println(findMinAndMax([]int{5,7,4,-2,-1,8}))
	// fmt.Println(findMinAndMax([]int{2,-5,-4,22,7,7}))
	// fmt.Println(findMinAndMax([]int{4,3,9,4,-21,7}))
	// fmt.Println(findMinAndMax([]int{-1,5,6,4,2,18}))
	// fmt.Println(findMinAndMax([]int{-2,5,-7,4,7,-20}))

	MaximumBuyProduct(50000, []int{25000,25000,10000,14000})
	MaximumBuyProduct(30000, []int{15000,10000,12000,5000,3000})
	MaximumBuyProduct(10000, []int{2000,3000,1000,2000,10000})
	MaximumBuyProduct(4000, []int{7500,3000,2500,3000})
	MaximumBuyProduct(0, []int{10000,30000})

}

func MaxSequence(data []int) int {
	var total int = 0
	if len(data)%2 == 0 {
		var mid int = len(data)/2 - 1

		for i := mid - 1; i <= mid+3; i++ {
			total += data[i]
		}

	} else {
		var mid int = len(data) / 2
		for i := mid - 1; i <= mid+2; i++ {
			total += data[i]
		}
	}
	return total

}

func primeX(number int) int {
	if number == 1 {
		return 2
	} else if number < 1 {
		return 0
	}

	var jumlah int = 1
	var i int = 1

	for jumlah <= number {
		if cekBilPrima(i) {
			jumlah++
		}
		i++
	}
	return i - 1

}

func cekBilPrima(number int) bool {
	var bagi int
	if number == 1 {
		return false
	} else if number%2 == 0 && number > 2 {
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

func cekStartPrime(start int) int {
	if cekBilPrima(start) {
		return start
	}
	return cekStartPrime(1 + start)

}

func primaSegiEmpat(high, wide, start int) {
	var total int

	for i := 0; i < wide; i++ {

		for j := 0; j < high; j++ {
			start = cekStartPrime(start + 1)
			total = total + start

			fmt.Print(start)
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Print(total)

	fmt.Print("\n\n")
}

func findMinAndMax(arr []int) string {
	var min int = arr[0]
	var max int = arr[0]

	var indexmin int
	var indexmax int

	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			indexmin = i
		} else if arr[i] > max {
			max = arr[i]
			indexmax = i
		}
	}
	return fmt.Sprintf("min: %d index: %d max: %d index: %d", min,indexmin,max,indexmax)
}

func MaximumBuyProduct(money int, productPrice []int) int {
	var i int
	var jmlmax int
	var res int
	sort.Ints(productPrice)

	for money >=0 && i <len(productPrice) {
		money -= productPrice[i]
		jmlmax++
		i++
	}
	res = jmlmax -1
	fmt.Println(res)
	return res
}