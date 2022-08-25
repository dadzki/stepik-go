package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	getSum(number)
}

func getSum(n int) {
	sum := 0
	for n > 0 {
		sum = sum + n%10
		n = n / 10
	}
	if sum < 10 {
		fmt.Println(sum)
	} else {
		getSum(sum)
	}
}
