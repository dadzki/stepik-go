package main

import "fmt"

func main() {
	var num, leftNumber, rightNumber uint

	fmt.Scan(&num)

	leftNumber = num/100000 + num/10000%10 + num/1000%10
	rightNumber = num/100%10 + num/10%10 + num%10

	if leftNumber == rightNumber {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
