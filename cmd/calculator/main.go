package main

import (
	"fmt"

	"learning-go/internal/calc"
)

func main() {
	var x int
	var y int

	fmt.Println("Hello Welcome to the CALC package")
	fmt.Println("Enter the first num: ")
	fmt.Scan(&x)
	fmt.Println("Enter the second num: ")
	fmt.Scan(&y)
	hasil := calc.Add(x, y)
	fmt.Printf("Hasilnya adalah: %v\n", hasil)
}
