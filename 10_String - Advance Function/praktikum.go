package main

import (
	"fmt"
	"strings"
)

// type Vehicle struct {
// 	Make string
// 	Name string
// }
// type Car struct {
// 	Vehicle
// }

// func (v *Vehicle) infor() {
// 	fmt.Println("make: ",v.Make)
// 	fmt.Println("make: ",v.Name)
// }

// func (c *Car) info() {
// 	fmt.Println("make: ", c.Make)
// 	fmt.Println("name: ", c.Name)
// }

// type Shape interface {
// 	getArea() int
// }
// type Rectangle struct {
// 	Length int
// 	Width int
// }

// func (r *Rectangle) getArea() int {
// 	return r.Length * r.Width
// }


// type BookService struct{}

// var storage []string = []string{}

// func (b *BookService)GetAll() []string{
// 	return storage
// }
// func (b *BookService)AddTitle(title string) []string {
// 	storage = append(storage,title)
// 	return storage
// }

func Compare(a, b string) string {
	if len(a) < len(b) {
		a,b = b,a

	}
	res := strings.Contains(a,b)
	if res {
		return b
		
	}
	return ""
}

// func caesar(offset int, input string) string {
	
// }


func main(){
	// fmt.Println(caesar(3, "abc")) //def
	// fmt.Println(caesar(2, "alta")) //cnvc
	// fmt.Println(caesar(10, "alterraacademy")) //kvdobbkkmknowi
	// fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) //bcdefghijklmnopqrstuvwxyza
	// fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) //mnopqrstuvwxyzabcdefghijkl


	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))

	// var shape Shape = &Rectangle{
	// 	getArea(Shape)
	// }
// fmt.Println(sumData(2,1,3))
// fmt.Println(convToCamelCase("Alterra Academy"))
// fmt.Println(anagramCheck("kasur","roesak"))

//tipe 2
// car := Car{
// 	Vehicle{
// 		Make : "Arema",
// 		Name : "Sasaji",
// 	},
// }
// car.info()

//tipe dasar
// car := Car{
// 	Make: "RUF",
// 	Name: "CTR Yellowbird",
// }

// car.info()


// var drinks []string = []string{"water","melon"}

// var drinkChannel chan string = make(chan string)

// go func (drinks []string) {
// 	for _, drink := range drinks {
// 		drinkChannel <- drink

// 	}
// 	close(drinkChannel)
// }(drinks)

// for drink := range drinkChannel{
// 	fmt.Println(drink)
// }

// }

// func sumData(data ...int) int {
// 	result := 0
// 	for _, val := range data {
// 		result = result + val
// 		fmt.Println("woi", result)
// 	}
	
// 	return result

// }

// func convToCamelCase(sample string) string {
// 	// var result string = ""
// 	low := strings.ToLower(sample)
// 	return low
// }

// func anagramCheck(kata1, kata2 string) bool {

// 	var freq1 map[string]int = map[string]int{}
// 	var freq2 map[string]int = map[string]int{}

// 	for _, value := range kata1 {
// 		freq1[string(value)]++
// 	}

// 	for _, value := range kata2 {
// 		freq2[string(value)]++
// 	}

// 	for _, value := range kata2 {
// 	if freq1[string(value)] != freq2[string(value)] {
// 		return false
// 	}
// }
// 	return true
} 



