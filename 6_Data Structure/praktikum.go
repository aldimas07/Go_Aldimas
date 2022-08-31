package main

// import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	slice := append(arrayA, arrayB...)
	return slice

}
func main() {
	// fmt.Println(munculSekali("1234123"))
	// fmt.Println(munculSekali("76523752"))
	// fmt.Println(munculSekali("12345"))
	// fmt.Println(munculSekali("1122334455"))
	// fmt.Println(munculSekali("0872504"))


	//Output Problem 1
	// fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// fmt.Println(ArrayMerge([]string{}, []string{"hworang"}))

	// fmt.Println(ArrayMerge([]string{}, []string{}))
}

// func munculSekali(angka string) []int{
// 	// var jml int
// 	for i:=0; i <= len(angka); i++{
// 		for j:=i+1; j < len(angka); j++{
		
// 		}
// 	}
// }
