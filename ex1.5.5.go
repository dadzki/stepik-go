package main

import "fmt"

func main() {

	var num uint

	fmt.Scan(&num)
	if num <= 10000 && num >= 0 {
		fmt.Println(num / 10 % 10)
	}

}
