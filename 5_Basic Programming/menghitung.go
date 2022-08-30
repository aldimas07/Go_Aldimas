package main

import "fmt"

func main() {
	// tabung()
	// gradenilai()
	// faktorbilangan()

	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))

	// menghitung luas permukaan tabung
	// var pi float64 = 3.14
	// var inputr float64
	// var inputt float64
	// var res float64

	// fmt.Println("Masukkan jari jari lalu enter dan masukkan tinggi tabung: ")
	// fmt.Scanf("%f\n%f", &inputr, &inputt)
	// res = (2*pi*inputr*inputr) + (2*pi*inputr*inputt)
	// fmt.Println(res)

}

func tabung() {
	//Problem 1
	var T float64 = 20
	var pi float64 = 3.14
	var r float64 = 4
	var res float64

	res = (2 * pi * r * r) + (2 * pi * r * T)
	fmt.Println(res)
}

func gradenilai() {
	//Problem 2
	var studentScore int
	var nama string

	fmt.Println("Masukkan nama siswa dan nilai siswa: ")
	fmt.Scanf("%s\n%d", &nama, &studentScore)

	if studentScore >= 80 && studentScore <= 100 {
		fmt.Println("Nama Siswa: ", nama)
		fmt.Println("Nilai Siswa: A")
	} else if studentScore >= 65 && studentScore <= 79 {
		fmt.Println("Nama Siswa: ", nama)
		fmt.Println("Nilai Siswa: B")
	} else if studentScore >= 50 && studentScore <= 64 {
		fmt.Println("Nama Siswa: ", nama)
		fmt.Println("Nilai Siswa: C")
	} else if studentScore >= 35 && studentScore <= 49 {
		fmt.Println("Nama Siswa: ", nama)
		fmt.Println("Nilai Siswa: D")
	} else if studentScore >= 0 && studentScore <= 34 {
		fmt.Println("Nama Siswa: ", nama)
		fmt.Println("Nilai Siswa: E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}

func faktorbilangan() {
	var n int
	fmt.Print("Input: ")
	fmt.Scanf("%d", &n)
	fmt.Println("Output: ")
	if n > 1 {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				fmt.Println(i)
			}
		}
	} else {
		fmt.Println("Angka harus lebih besar dari 1")
	}

}

func primeNumber(number int) bool {
	var bagi int

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			bagi = bagi + 1
		}
	}

	if bagi == 2 {
		fmt.Print("Bilangan Prima\n")
		return true

	} else {
		fmt.Print("Bukan Bilangan Prima\n")
		return false

	}

}
