package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	slice := append(arrayA, arrayB...)
	return slice

}
func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"hworang"}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}
