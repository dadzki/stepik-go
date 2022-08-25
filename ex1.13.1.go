package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)

	sum = n/100 + n/10%10 + n%10

	fmt.Println(sum)
}
