package main

import "fmt"

func main() {

	var num int

	fmt.Scan(&num)
	if num <= 10000 {
		fmt.Println(num % 100)
	}

}
