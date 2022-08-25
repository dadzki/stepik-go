package main

import "fmt"

func main() {
	var num, max, count int
	num = 1
	max = 0
	for num > 0 {
		fmt.Scan(&num)

		if num > max {
			max = num
			count = 0
		}
		if num == max {
			count++
		}
	}
	fmt.Println(count)
}
