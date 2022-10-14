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
		a, b = b, a

	}
	res := strings.Contains(a, b)
	if res {
		return b

	}
	return ""
}

func caesar(offset int, input string) string {
	res := make([]byte, 0, len(input))

	for _, v := range input {
		conv := (int(v)+offset-97)%26 + 97
		res = append(res, byte(conv))
	}
	return string(res)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func getMinMax(numbers ...*int) (min int, max int) {

	min = *numbers[0]
	max = *numbers[0]

	for _, value := range numbers {
		if *value < min {
			min = *value
		} else if *value > max {
			max = *value
		}
	}
	return min, max
}

// type Student struct {
// 	Name  []string
// 	Score []int
// }

// func (s Student) Average() float64 {
// 	var average float64
// 	for _, score := range s.Score {
// 		average += float64(score)
// 	}
// 	average = average / float64(len(s.Score))
// 	fmt.Println(float64(len(s.Score)))
// 	return average

// }

// func (s Student) Min() (min int, name string) {
// 	min = 100

// 	var index int
// 	for i, score := range s.Score {
// 		if score < min {
// 			min = score
// 			index = i
// 		}
// 	}
// 	return s.Score[index], s.Name[index]

// }
// func (s Student) Max() (max int, name string) {
// 	max = 0

// 	var index int
// 	for i, score := range s.Score {
// 		if score > max {
// 			max = score
// 			index = i
// 		}
// 	}
// 	return s.Score[index], s.Name[index]
// }

// Problem 6 - Substitution Cipher
type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	result := make([]string, 0, len(s.name))

	for _, char := range s.name {
		if char >= 97 && char <= 122 {
			char = (char - 97 + 13) % 26 + 97
		}
		result = append(result, string(char))
	}
	nameEncode = strings.Join(result, "")

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""
	result := make([]string, 0, len(s.nameEncode))
	for _, char := range s.nameEncode {
		if char >= 'a' || char <= 'z' {
			char = (char - 97 +39) % 26 + 97
		}
		result = append(result, string(char))
	}
	nameDecode = strings.Join(result, "")
	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a
	fmt.Print("[1]Encrypt \n[2]Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student's name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}

	// var a = Student{}

	// for i := 0; i < 6; i++ {
	// 	var name string
	// 	fmt.Print("Input" + string(i) + "Student's Name : ")
	// 	fmt.Scan(&name)
	// 	a.Name = append(a.Name, name)
	// 	var score int
	// 	fmt.Print("Input" + name + "Score : ")
	// 	fmt.Scan(&score)
	// 	a.Score = append(a.Score, score)
	// }

	// fmt.Println("\n\nAverage Score Students is", a.Average())
	// scoreMax, nameMax := a.Max()
	// fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	// scoreMin, nameMin := a.Min()
	// fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")

	// var a1, a2, a3, a4, a5, a6, min, max int
	// fmt.Scan(&a1)
	// fmt.Scan(&a2)
	// fmt.Scan(&a3)
	// fmt.Scan(&a4)
	// fmt.Scan(&a5)
	// fmt.Scan(&a6)
	// min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	// fmt.Println(max, "is the maximum number")
	// fmt.Println(min, "is the minimum number")

	// a := 10
	// b := 20

	// swap(&a, &b)
	// fmt.Println(a, b)

	// fmt.Println(caesar(3, "abc")) //def
	// fmt.Println(caesar(2, "alta")) //cnvc
	// fmt.Println(caesar(10, "alterraacademy")) //kvdobbkkmknowi
	// fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) //bcdefghijklmnopqrstuvwxyza
	// fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) //mnopqrstuvwxyzabcdefghijkl

	// fmt.Println(Compare("AKA", "AKASHI"))
	// fmt.Println(Compare("KANGOORO", "KANG"))
	// fmt.Println(Compare("KI", "KIJANG"))
	// fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	// fmt.Println(Compare("ILALANG", "ILA"))

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
