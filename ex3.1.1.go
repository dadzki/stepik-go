package main

import (
	"fmt"
)

func work(x int) int {
	return x * x
}

func main() {
	var num int
	cache := make(map[int]int)
	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("ошибка ввода")
		}

		if _, ok := cache[num]; !ok {
			cache[num] = work(num)
		}

		fmt.Printf("%v ", cache[num])
	}
}
