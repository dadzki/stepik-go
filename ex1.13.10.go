package main

import "fmt"

func main() {
	var n, tmp int
	fmt.Scan(&n)

	tmp = n % 10
	switch {
	case (n >= 11 && n <= 14) || tmp == 0 || (tmp >= 5 && tmp <= 9):
		fmt.Println(n, "korov")
	case tmp == 1:
		fmt.Println(n, "korova")
	case tmp >= 2 && tmp <= 4:
		fmt.Println(n, "korovy")
	}
}
