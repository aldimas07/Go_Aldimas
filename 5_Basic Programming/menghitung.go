package main

import "fmt"
func main(){
	// tabung()
	var pi float64 = 3.14
	var inputr float64
	var inputt float64
	var res float64

	fmt.Println("Masukkan jari jari lalu enter dan masukkan tinggi tabung: ")
	fmt.Scanf("%f\n%f", &inputr, &inputt)
	res = (2*pi*inputr*inputr) + (2*pi*inputr*inputt)
	fmt.Println(res)

}

func tabung(){
	//Problem 1
	var T float64 = 20
	var pi float64 = 3.14
	var r float64 = 4
	var res float64

	res = (2*pi*r*r) + (2*pi*r*T)
	fmt.Println(res)
}
