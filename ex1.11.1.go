package main

import "fmt"

func main() {
	var number1 float64
	fmt.Scan(&number1)

	if number1 > 10000 {
		fmt.Printf("%e", number1)
	} else if number1 <= 0 {
		fmt.Printf("число %2.2f не подходит", number1)
	} else {
		fmt.Printf("%.4f", number1*number1)
	}
}
