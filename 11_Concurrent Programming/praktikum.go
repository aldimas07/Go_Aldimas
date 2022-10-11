package main

import (
	"fmt"
	"time"
)

func parallel(char string) {
	for i := 0; i < len(char); i++ {
		res := char[i]
		fmt.Printf("%c : %d \n", res, res)
	}
}
func main() {
	go parallel("lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua")
	time.Sleep(1 * time.Second)
}
