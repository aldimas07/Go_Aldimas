package main

import (
	"fmt"
	"strconv"
	// "golang.org/x/text/number"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	var res []string = []string{}

	var storage map[string]bool = map[string]bool{}

	for _, item1 := range arrayA {
		res = append(res, item1)
		storage[item1] = true
	}

	for _, item2 := range arrayB {
		_, pencarian := storage[item2] //pengecekan

		if !pencarian  { //jika value yang di cari tidak ada maka akan ditambahkan ke slice res
			res = append(res, item2) 
		}
	}
	return res
}
func main() {
	
	// Output Problem 1
	// fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	
	// fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	
	// fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	
	// fmt.Println(ArrayMerge([]string{}, []string{"hworang"}))
	
	// fmt.Println(ArrayMerge([]string{}, []string{}))
	
	// Output Problem 2
	// fmt.Println(munculSekali("1234123"))
	// fmt.Println(munculSekali("76523752"))
	// fmt.Println(munculSekali("12345"))
	// fmt.Println(munculSekali("1122334455"))
	// fmt.Println(munculSekali("0872504"))

	// Output Problem 3

	fmt.Println(PairSum([]int{1,2,3,4,6},6))
	fmt.Println(PairSum([]int{2,5,9,11},11))
	fmt.Println(PairSum([]int{1,3,5,7},12))
	fmt.Println(PairSum([]int{1,4,6,8},10))
	fmt.Println(PairSum([]int{1,5,6,7},6))

}

func munculSekali(angka string) []int {
	var hasil []int = []int{}

	var freq map[string]int = map[string]int{}

	for _, num := range angka {
		freq[string(num)]++
	}
	
	for key, frekuensi := range freq {
		if frekuensi == 1{
			res, _ := strconv.Atoi(key)
			hasil = append(hasil, res)
		}
	}
	return hasil
}

func PairSum(arr []int, target int) []int {
		var res []int = []int{}
		var storage map[int]int = map[int]int{}

		for i, y := range arr {
			var x int = target - y


			_, pencarian := storage[x]
			if pencarian {
				res = append(res, storage[x], i)
			}
			storage[y] = i
		}
		return res
}
